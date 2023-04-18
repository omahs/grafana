// Code generated by mockery v2.24.0. DO NOT EDIT.

package dashboards

import (
	context "context"

	folder "github.com/grafana/grafana/pkg/services/folder"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/services/alerting/models"

	quota "github.com/grafana/grafana/pkg/services/quota"
)

// FakeDashboardStore is an autogenerated mock type for the Store type
type FakeDashboardStore struct {
	mock.Mock
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *FakeDashboardStore) Count(_a0 context.Context, _a1 *quota.ScopeParameters) (*quota.Map, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *quota.Map
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *quota.ScopeParameters) (*quota.Map, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *quota.ScopeParameters) *quota.Map); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*quota.Map)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *quota.ScopeParameters) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDashboardsInFolder provides a mock function with given fields: ctx, request
func (_m *FakeDashboardStore) CountDashboardsInFolder(ctx context.Context, request *CountDashboardsInFolderRequest) (int64, error) {
	ret := _m.Called(ctx, request)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CountDashboardsInFolderRequest) (int64, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CountDashboardsInFolderRequest) int64); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CountDashboardsInFolderRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteACLByUser provides a mock function with given fields: _a0, _a1
func (_m *FakeDashboardStore) DeleteACLByUser(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDashboard provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) DeleteDashboard(ctx context.Context, cmd *DeleteDashboardCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteDashboardCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDashboardsInFolder provides a mock function with given fields: ctx, request
func (_m *FakeDashboardStore) DeleteDashboardsInFolder(ctx context.Context, request *DeleteDashboardsInFolderRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteDashboardsInFolderRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedProvisionedDashboards provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) DeleteOrphanedProvisionedDashboards(ctx context.Context, cmd *DeleteOrphanedProvisionedDashboardsCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteOrphanedProvisionedDashboardsCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) FindDashboards(ctx context.Context, query *FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
	ret := _m.Called(ctx, query)

	var r0 []DashboardSearchProjection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) []DashboardSearchProjection); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DashboardSearchProjection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *FindPersistedDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboard(ctx context.Context, query *GetDashboardQuery) (*Dashboard, error) {
	ret := _m.Called(ctx, query)

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) (*Dashboard, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) *Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardACLInfoList provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardACLInfoList(ctx context.Context, query *GetDashboardACLInfoListQuery) ([]*DashboardACLInfoDTO, error) {
	ret := _m.Called(ctx, query)

	var r0 []*DashboardACLInfoDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardACLInfoListQuery) ([]*DashboardACLInfoDTO, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardACLInfoListQuery) []*DashboardACLInfoDTO); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*DashboardACLInfoDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardACLInfoListQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardTags provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardTags(ctx context.Context, query *GetDashboardTagsQuery) ([]*DashboardTagCloudItem, error) {
	ret := _m.Called(ctx, query)

	var r0 []*DashboardTagCloudItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) ([]*DashboardTagCloudItem, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) []*DashboardTagCloudItem); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*DashboardTagCloudItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardTagsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardUIDByID provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardUIDByID(ctx context.Context, query *GetDashboardRefByIDQuery) (*DashboardRef, error) {
	ret := _m.Called(ctx, query)

	var r0 *DashboardRef
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIDQuery) (*DashboardRef, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIDQuery) *DashboardRef); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardRef)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardRefByIDQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboards(ctx context.Context, query *GetDashboardsQuery) ([]*Dashboard, error) {
	ret := _m.Called(ctx, query)

	var r0 []*Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) ([]*Dashboard, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) []*Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardsByPluginID provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardsByPluginID(ctx context.Context, query *GetDashboardsByPluginIDQuery) ([]*Dashboard, error) {
	ret := _m.Called(ctx, query)

	var r0 []*Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsByPluginIDQuery) ([]*Dashboard, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsByPluginIDQuery) []*Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardsByPluginIDQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDashboardData provides a mock function with given fields: ctx, name
func (_m *FakeDashboardStore) GetProvisionedDashboardData(ctx context.Context, name string) ([]*DashboardProvisioning, error) {
	ret := _m.Called(ctx, name)

	var r0 []*DashboardProvisioning
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*DashboardProvisioning, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*DashboardProvisioning); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*DashboardProvisioning)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardID provides a mock function with given fields: ctx, dashboardID
func (_m *FakeDashboardStore) GetProvisionedDataByDashboardID(ctx context.Context, dashboardID int64) (*DashboardProvisioning, error) {
	ret := _m.Called(ctx, dashboardID)

	var r0 *DashboardProvisioning
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*DashboardProvisioning, error)); ok {
		return rf(ctx, dashboardID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *DashboardProvisioning); ok {
		r0 = rf(ctx, dashboardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardProvisioning)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, dashboardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardUID provides a mock function with given fields: ctx, orgID, dashboardUID
func (_m *FakeDashboardStore) GetProvisionedDataByDashboardUID(ctx context.Context, orgID int64, dashboardUID string) (*DashboardProvisioning, error) {
	ret := _m.Called(ctx, orgID, dashboardUID)

	var r0 *DashboardProvisioning
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*DashboardProvisioning, error)); ok {
		return rf(ctx, orgID, dashboardUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *DashboardProvisioning); ok {
		r0 = rf(ctx, orgID, dashboardUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardProvisioning)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, dashboardUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAdminPermissionInDashboardsOrFolders provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) HasAdminPermissionInDashboardsOrFolders(ctx context.Context, query *folder.HasAdminPermissionInDashboardsOrFoldersQuery) (bool, error) {
	ret := _m.Called(ctx, query)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *folder.HasAdminPermissionInDashboardsOrFoldersQuery) (bool, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *folder.HasAdminPermissionInDashboardsOrFoldersQuery) bool); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *folder.HasAdminPermissionInDashboardsOrFoldersQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasEditPermissionInFolders provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) HasEditPermissionInFolders(ctx context.Context, query *folder.HasEditPermissionInFoldersQuery) (bool, error) {
	ret := _m.Called(ctx, query)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *folder.HasEditPermissionInFoldersQuery) (bool, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *folder.HasEditPermissionInFoldersQuery) bool); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *folder.HasEditPermissionInFoldersQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAlerts provides a mock function with given fields: ctx, dashID, alerts
func (_m *FakeDashboardStore) SaveAlerts(ctx context.Context, dashID int64, alerts []*models.Alert) error {
	ret := _m.Called(ctx, dashID, alerts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.Alert) error); ok {
		r0 = rf(ctx, dashID, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) SaveDashboard(ctx context.Context, cmd SaveDashboardCommand) (*Dashboard, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, SaveDashboardCommand) (*Dashboard, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, SaveDashboardCommand) *Dashboard); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, SaveDashboardCommand) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveDashboardWithMetadata provides a mock function with given fields: ctx, msg, dash, provisioning
func (_m *FakeDashboardStore) SaveDashboardWithMetadata(ctx context.Context, msg string, dash *Dashboard, provisioning *DashboardProvisioning) (*Dashboard, error) {
	ret := _m.Called(ctx, msg, dash, provisioning)

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *Dashboard, *DashboardProvisioning) (*Dashboard, error)); ok {
		return rf(ctx, msg, dash, provisioning)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *Dashboard, *DashboardProvisioning) *Dashboard); ok {
		r0 = rf(ctx, msg, dash, provisioning)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *Dashboard, *DashboardProvisioning) error); ok {
		r1 = rf(ctx, msg, dash, provisioning)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveProvisionedDashboard provides a mock function with given fields: ctx, cmd, provisioning
func (_m *FakeDashboardStore) SaveProvisionedDashboard(ctx context.Context, cmd SaveDashboardCommand, provisioning *DashboardProvisioning) (*Dashboard, error) {
	ret := _m.Called(ctx, cmd, provisioning)

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, SaveDashboardCommand, *DashboardProvisioning) (*Dashboard, error)); ok {
		return rf(ctx, cmd, provisioning)
	}
	if rf, ok := ret.Get(0).(func(context.Context, SaveDashboardCommand, *DashboardProvisioning) *Dashboard); ok {
		r0 = rf(ctx, cmd, provisioning)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, SaveDashboardCommand, *DashboardProvisioning) error); ok {
		r1 = rf(ctx, cmd, provisioning)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnprovisionDashboard provides a mock function with given fields: ctx, id
func (_m *FakeDashboardStore) UnprovisionDashboard(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDashboardACL provides a mock function with given fields: ctx, uid, items
func (_m *FakeDashboardStore) UpdateDashboardACL(ctx context.Context, uid int64, items []*DashboardACL) error {
	ret := _m.Called(ctx, uid, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*DashboardACL) error); ok {
		r0 = rf(ctx, uid, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDashboardBeforeSave provides a mock function with given fields: ctx, dashboard, overwrite
func (_m *FakeDashboardStore) ValidateDashboardBeforeSave(ctx context.Context, dashboard *Dashboard, overwrite bool) (bool, error) {
	ret := _m.Called(ctx, dashboard, overwrite)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Dashboard, bool) (bool, error)); ok {
		return rf(ctx, dashboard, overwrite)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Dashboard, bool) bool); ok {
		r0 = rf(ctx, dashboard, overwrite)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Dashboard, bool) error); ok {
		r1 = rf(ctx, dashboard, overwrite)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFakeDashboardStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeDashboardStore creates a new instance of FakeDashboardStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeDashboardStore(t mockConstructorTestingTNewFakeDashboardStore) *FakeDashboardStore {
	mock := &FakeDashboardStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
