package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"time"

	dtos "github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/models/roletype"
	contextmodel "github.com/grafana/grafana/pkg/services/contexthandler/model"
	"github.com/grafana/grafana/pkg/services/dashboards"
	"github.com/grafana/grafana/pkg/services/rendering"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/web"
)

func (hs *HTTPServer) GetSlackChannels(c *contextmodel.ReqContext) response.Response {
	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/conversations.list", nil)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hs.Cfg.SlackToken))

	if err != nil {
		return response.Error(http.StatusInternalServerError, "could not create request", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response.Error(http.StatusInternalServerError, "error making http request", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			hs.log.Error("failed to close response body", "err", err)
		}
	}()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.Error(http.StatusInternalServerError, "could not read response body", err)
	}

	var result dtos.SlackChannels
	err = json.Unmarshal(b, &result)
	if err != nil {
		return response.Error(http.StatusInternalServerError, "could not unmarshall response", err)
	}

	return response.JSON(http.StatusOK, result.Channels)
}

func (hs *HTTPServer) ShareToSlack(c *contextmodel.ReqContext) response.Response {
	dashboardUid := web.Params(c.Req)[":uid"]

	dashboard, err := hs.DashboardService.GetDashboard(c.Req.Context(), &dashboards.GetDashboardQuery{UID: dashboardUid})
	if err != nil {
		hs.log.Error("fail to get dashboard", "err", err, "dashboard UID", dashboardUid)
		return response.Error(400, "error parsing body", err)
	}

	var shareRequest ShareRequest
	if err := web.Bind(c.Req, &shareRequest); err != nil {
		return response.Error(400, "error parsing body", err)
	}

	protocol := hs.Cfg.Protocol
	switch protocol {
	case setting.HTTPScheme:
		protocol = "http"
	case setting.HTTP2Scheme, setting.HTTPSScheme:
		protocol = "https"
	default:
		// TODO: Handle other schemes?
	}

	domain := "localhost"
	if hs.Cfg.HTTPAddr != "0.0.0.0" {
		domain = hs.Cfg.HTTPAddr
	}

	dashboardLink := fmt.Sprintf("%s://%s:%s%s", protocol, domain, hs.Cfg.HTTPPort, shareRequest.DashboardPath)

	blocks := []Block{
		{
			Type: "section",
			Text: &Text{
				Type: "mrkdwn",
				Text: fmt.Sprintf("<%s|*%s*>", dashboardLink, dashboard.Title),
			},
		},
		{
			Type: "section",
			Text: &Text{
				Type: "plain_text",
				Text: shareRequest.Message,
			},
		},
		{
			Type: "image",
			Title: &Text{
				Type: "plain_text",
				Text: "Dashboard preview",
			},
			ImageURL: shareRequest.ImagePreviewUrl,
			AltText:  "dashboard preview",
		},
		{
			Type: "actions",
			Elements: []Element{{
				Type: "button",
				Text: &Text{
					Type: "plain_text",
					Text: "View in Grafana",
				},
				Style: "primary",
				Value: "View in Grafana",
				URL:   dashboardLink,
			}},
		},
	}

	for _, channelId := range shareRequest.ChannelIds {
		postMessageRequest := &PostMessageRequest{
			Channel: channelId,
			Blocks:  blocks,
		}

		jsonBody, err := json.Marshal(postMessageRequest)

		if err != nil {
			return response.Error(400, "error parsing body", err)
		}

		req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", bytes.NewReader(jsonBody))
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hs.Cfg.SlackToken))

		if err != nil {
			fmt.Errorf("client: could not create request: %w", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Errorf("client: error making http request: %w", err)
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Errorf("client: could not read response body: %w", err)
		}

		hs.log.Info("successfully sent unfurl event payload", "body", string(resBody))

	}

	return response.JSON(http.StatusOK, nil)
}

func (hs *HTTPServer) AcknowledgeSlackEvent(c *contextmodel.ReqContext) response.Response {
	var eventPayload EventPayload
	if err := web.Bind(c.Req, &eventPayload); err != nil {
		return response.Error(http.StatusBadRequest, "error parsing body", err)
	}

	switch eventPayload.Type {
	case "url_verification":
		return response.JSON(http.StatusOK, &EventChallengeAck{
			Challenge: eventPayload.Challenge,
		})
	case "event_callback":
		if eventPayload.Event.Type == "link_shared" && eventPayload.Event.Source == "conversations_history" {
			defer hs.handleLinkSharedEvent(eventPayload)
			return response.Empty(http.StatusOK)
		}
	}

	return response.Error(http.StatusBadRequest, "not handling this event type", fmt.Errorf("not handling this event type"))
}

func (hs *HTTPServer) handleLinkSharedEvent(event EventPayload) {
	ctx := context.Background()

	// TODO: handle multiple links
	renderPath, dashboardUID := extractURLInfo(event.Event.Links[0].URL)
	if renderPath == "" {
		hs.log.Error("fail to extract render path from link")
		return
	}

	dashboard, err := hs.DashboardService.GetDashboard(ctx, &dashboards.GetDashboardQuery{UID: dashboardUID})
	if err != nil {
		hs.log.Error("fail to get dashboard", "err", err, "dashboard UID", dashboardUID)
		return
	}

	imagePath, err := hs.renderDashboard(ctx, renderPath)
	if err != nil {
		hs.log.Error("fail to render dashboard for Slack preview", "err", err)
		return
	}

	err = hs.sendUnfurlEvent(ctx, event, imagePath, dashboard.Title)
	if err != nil {
		hs.log.Error("fail to send unfurl event to Slack", "err", err)
	}
}

func (hs *HTTPServer) sendUnfurlEvent(c context.Context, linkEvent EventPayload, imagePath string, dashboardTitle string) error {
	eventPayload := &UnfurlEventPayload{
		Channel: linkEvent.Event.Channel,
		TS:      linkEvent.Event.MessageTS,
		Unfurls: make(Unfurls),
	}

	imageFileName := filepath.Base(imagePath)
	imageURL := hs.getImageURL(imageFileName)
	for _, link := range linkEvent.Event.Links {
		eventPayload.Unfurls[link.URL] = Unfurl{
			Blocks: []Block{
				{
					Type: "header",
					Text: &Text{
						Type: "plain_text",
						Text: dashboardTitle,
					},
				},
				{
					Type: "section",
					Text: &Text{
						Type: "plain_text",
						Text: "Here is the dashboard that I wanted to show you",
					},
				},
				{
					Type: "image",
					Title: &Text{
						Type: "plain_text",
						Text: "Dashboard preview",
					},
					ImageURL: imageURL,
					AltText:  "dashboard preview",
				},
				{
					Type: "actions",
					Elements: []Element{{
						Type: "button",
						Text: &Text{
							Type: "plain_text",
							Text: "View Dashboard",
						},
						Style: "primary",
						Value: link.URL,
						URL:   link.URL,
					}},
				},
			},
		}
	}

	b, err := json.Marshal(eventPayload)
	if err != nil {
		return fmt.Errorf("client: could not create body: %w", err)
	}
	hs.log.Info("Posting to slack api", "eventPayload", string(b))

	bodyReader := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.unfurl", bodyReader)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hs.Cfg.SlackToken))
	if err != nil {
		return fmt.Errorf("client: could not create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %w", err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			hs.log.Error("failed to close response body", "err", err)
		}
	}()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("client: could not read response body: %w", err)
	}

	hs.log.Info("successfully sent unfurl event payload", "body", string(resBody))
	return nil
}

// TODO: Duplicated from the rendering service - maybe we can do this in another way to not duplicate this
func (hs *HTTPServer) getImageURL(imageName string) string {
	if hs.Cfg.RendererCallbackUrl != "" {
		return fmt.Sprintf("%s%s/%s", hs.Cfg.RendererCallbackUrl, "public/img/attachments", imageName)
	}

	protocol := hs.Cfg.Protocol
	switch protocol {
	case setting.HTTPScheme:
		protocol = "http"
	case setting.HTTP2Scheme, setting.HTTPSScheme:
		protocol = "https"
	default:
		// TODO: Handle other schemes?
	}

	subPath := ""
	if hs.Cfg.ServeFromSubPath {
		subPath = hs.Cfg.AppSubURL
	}

	domain := "localhost"
	if hs.Cfg.HTTPAddr != "0.0.0.0" {
		domain = hs.Cfg.HTTPAddr
	}

	return fmt.Sprintf("%s://%s:%s%s/%s/%s", protocol, domain, hs.Cfg.HTTPPort, subPath, "public/img/attachments", imageName)
}

// extractURLInfo returns the render path and the dashboard UID
func extractURLInfo(dashboardURL string) (string, string) {
	re := regexp.MustCompile(`.*(\/d\/([^\/]*)\/.*)`)
	res := re.FindStringSubmatch(dashboardURL)
	if len(res) != 3 {
		return "", ""
	}

	return res[1], res[2]
}

func (hs *HTTPServer) renderDashboard(ctx context.Context, renderPath string) (string, error) {
	result, err := hs.RenderService.Render(ctx, rendering.Opts{
		TimeoutOpts: rendering.TimeoutOpts{
			Timeout: time.Duration(60) * time.Second,
		},
		AuthOpts: rendering.AuthOpts{
			// TODO: get the org id from the URL
			OrgID:   1,
			OrgRole: roletype.RoleAdmin,
		},
		Width:  1600,
		Height: 800,
		//Path:   web.Params(c.Req)["*"] + queryParams,
		Path: renderPath,
		//Timezone:          queryReader.Get("tz", ""),
		//Encoding:          queryReader.Get("encoding", ""),
		ConcurrentLimit:   hs.Cfg.RendererConcurrentRequestLimit,
		DeviceScaleFactor: 1, // negative numbers will render larger and then scale down
		Theme:             models.ThemeDark,
	}, nil)
	if err != nil {
		return "", err
	}

	return result.FilePath, nil
}

func (hs *HTTPServer) GeneratePreview(c *contextmodel.ReqContext) response.Response {
	var previewRequest PreviewRequest
	if err := web.Bind(c.Req, &previewRequest); err != nil {
		return response.Error(http.StatusBadRequest, "error parsing body", err)
	}

	hs.log.Info("Generating preview", "dashboard_url", previewRequest.DashboardURL)

	filePath, err := hs.renderDashboard(c.Req.Context(), previewRequest.DashboardURL)
	if err != nil {
		return response.Error(http.StatusInternalServerError, "Rendering failed", err)
	}

	imageFileName := filepath.Base(filePath)
	imageURL := hs.getImageURL(imageFileName)

	return response.JSON(http.StatusOK, &PreviewResponse{
		PreviewURL: imageURL,
	})
}

type PreviewRequest struct {
	DashboardURL string `json:"dashboardUrl"`
}
type PreviewResponse struct {
	PreviewURL string `json:"previewUrl"`
}

type EventChallengeAck struct {
	Challenge string `json:"challenge"`
}

type EventPayload struct {
	Token              string          `json:"token"`
	TeamID             string          `json:"team_id"`
	APIAppID           string          `json:"api_app_id"`
	Event              Event           `json:"event"`
	Type               string          `json:"type"`
	EventID            string          `json:"event_id"`
	EventTime          int64           `json:"event_time"`
	Authorizations     []Authorization `json:"authorizations"`
	IsExtSharedChannel bool            `json:"is_ext_shared_channel"`
	EventContext       string          `json:"event_context"`
	Challenge          string          `json:"challenge"`
}

// Event represents the "event" field in the payload
type Event struct {
	Type            string `json:"type"`
	User            string `json:"user"`
	Channel         string `json:"channel"`
	MessageTS       string `json:"message_ts"`
	Links           []Link `json:"links"`
	Source          string `json:"source"`
	UnfurlID        string `json:"unfurl_id"`
	IsBotUserMember bool   `json:"is_bot_user_member"`
	EventTS         string `json:"event_ts"`
}

// Link represents the "links" field in the event
type Link struct {
	URL    string `json:"url"`
	Domain string `json:"domain"`
}

type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type ImageAccessory struct {
	Type     string `json:"type,omitempty"`
	Title    *Text  `json:"title,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}

type Element struct {
	Type  string `json:"type,omitempty"`
	Text  *Text  `json:"text,omitempty"`
	Style string `json:"style,omitempty"`
	Value string `json:"value,omitempty"`
	URL   string `json:"url,omitempty"`
}

type Block struct {
	Type      string          `json:"type,omitempty"`
	Text      *Text           `json:"text,omitempty"`
	Accessory *ImageAccessory `json:"accessory,omitempty"`
	ImageURL  string          `json:"image_url,omitempty"`
	Title     *Text           `json:"title,omitempty"`
	AltText   string          `json:"alt_text,omitempty"`
	Elements  []Element       `json:"elements,omitempty"`
}

type Unfurl struct {
	Blocks []Block `json:"blocks,omitempty"`
}

type Unfurls map[string]Unfurl

type UnfurlEventPayload struct {
	//Source   string  `json:"source"`
	//UnfurlID string  `json:"unfurl_id"`
	//Token    string  `json:"token"`
	Channel string  `json:"channel,omitempty"`
	TS      string  `json:"ts,omitempty"`
	Unfurls Unfurls `json:"unfurls,omitempty"`
}

type ShareRequest struct {
	ChannelIds      []string `json:"channelIds"`
	Message         string   `json:"message,omitempty"`
	ImagePreviewUrl string   `json:"imagePreviewUrl"`
	PanelId         string   `json:"panelId,omitempty"`
	DashboardPath   string   `json:"dashboardPath"`
}

type PostMessageRequest struct {
	Channel string  `json:"channel"`
	Blocks  []Block `json:"blocks,omitempty"`
}
