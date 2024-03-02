// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/iface/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	iface "boyi/pkg/iface"
	db "boyi/pkg/infra/db"
	dto "boyi/pkg/model/dto"
	option "boyi/pkg/model/option"
	context "context"
	sql "database/sql"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
	gorm "gorm.io/gorm"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// BatchInsert mocks base method.
func (m *MockIRepository) BatchInsert(ctx context.Context, tx *gorm.DB, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchInsert", ctx, tx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchInsert indicates an expected call of BatchInsert.
func (mr *MockIRepositoryMockRecorder) BatchInsert(ctx, tx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsert", reflect.TypeOf((*MockIRepository)(nil).BatchInsert), ctx, tx, data)
}

// Count mocks base method.
func (m *MockIRepository) Count(ctx context.Context, tx *gorm.DB, opt iface.WhereOption) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, tx, opt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIRepositoryMockRecorder) Count(ctx, tx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIRepository)(nil).Count), ctx, tx, opt)
}

// Create mocks base method.
func (m *MockIRepository) Create(ctx context.Context, tx *gorm.DB, data iface.Model, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, data}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIRepositoryMockRecorder) Create(ctx, tx, data interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, data}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRepository)(nil).Create), varargs...)
}

// CreateIfNotExists mocks base method.
func (m *MockIRepository) CreateIfNotExists(ctx context.Context, tx *gorm.DB, data iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, data, opt}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIfNotExists", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIfNotExists indicates an expected call of CreateIfNotExists.
func (mr *MockIRepositoryMockRecorder) CreateIfNotExists(ctx, tx, data, opt interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, data, opt}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIfNotExists", reflect.TypeOf((*MockIRepository)(nil).CreateIfNotExists), varargs...)
}

// CreateOrUpdate mocks base method.
func (m *MockIRepository) CreateOrUpdate(ctx context.Context, tx *gorm.DB, data iface.Model, opt iface.WhereOption, updateCol iface.UpdateColumns) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, tx, data, opt, updateCol)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockIRepositoryMockRecorder) CreateOrUpdate(ctx, tx, data, opt, updateCol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockIRepository)(nil).CreateOrUpdate), ctx, tx, data, opt, updateCol)
}

// Delete mocks base method.
func (m *MockIRepository) Delete(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, model, opt}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRepositoryMockRecorder) Delete(ctx, tx, model, opt interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, model, opt}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRepository)(nil).Delete), varargs...)
}

// DeleteMerchantDB mocks base method.
func (m *MockIRepository) DeleteMerchantDB(ctx context.Context, merchantId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMerchantDB", ctx, merchantId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMerchantDB indicates an expected call of DeleteMerchantDB.
func (mr *MockIRepositoryMockRecorder) DeleteMerchantDB(ctx, merchantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMerchantDB", reflect.TypeOf((*MockIRepository)(nil).DeleteMerchantDB), ctx, merchantId)
}

// Get mocks base method.
func (m *MockIRepository) Get(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, model, opt}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIRepositoryMockRecorder) Get(ctx, tx, model, opt interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, model, opt}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRepository)(nil).Get), varargs...)
}

// GetALLMerchantDB mocks base method.
func (m *MockIRepository) GetALLMerchantDB(ctx context.Context) (map[uint64]*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetALLMerchantDB", ctx)
	ret0, _ := ret[0].(map[uint64]*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetALLMerchantDB indicates an expected call of GetALLMerchantDB.
func (mr *MockIRepositoryMockRecorder) GetALLMerchantDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetALLMerchantDB", reflect.TypeOf((*MockIRepository)(nil).GetALLMerchantDB), ctx)
}

// GetDB mocks base method.
func (m *MockIRepository) GetDB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// GetDB indicates an expected call of GetDB.
func (mr *MockIRepositoryMockRecorder) GetDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockIRepository)(nil).GetDB))
}

// GetLast mocks base method.
func (m *MockIRepository) GetLast(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, model, opt}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLast", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLast indicates an expected call of GetLast.
func (mr *MockIRepositoryMockRecorder) GetLast(ctx, tx, model, opt interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, model, opt}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLast", reflect.TypeOf((*MockIRepository)(nil).GetLast), varargs...)
}

// GetMerchantDB mocks base method.
func (m *MockIRepository) GetMerchantDB(ctx context.Context, merchantId uint64) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerchantDB", ctx, merchantId)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMerchantDB indicates an expected call of GetMerchantDB.
func (mr *MockIRepositoryMockRecorder) GetMerchantDB(ctx, merchantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerchantDB", reflect.TypeOf((*MockIRepository)(nil).GetMerchantDB), ctx, merchantId)
}

// GetUserByID mocks base method.
func (m *MockIRepository) GetUserByID(ctx context.Context, userID uint64) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIRepositoryMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIRepository)(nil).GetUserByID), ctx, userID)
}

// GetUserIDs mocks base method.
func (m *MockIRepository) GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDs", ctx, opt)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDs indicates an expected call of GetUserIDs.
func (mr *MockIRepositoryMockRecorder) GetUserIDs(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDs", reflect.TypeOf((*MockIRepository)(nil).GetUserIDs), ctx, opt)
}

// List mocks base method.
func (m *MockIRepository) List(ctx context.Context, tx *gorm.DB, data interface{}, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, data, opt}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIRepositoryMockRecorder) List(ctx, tx, data, opt interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, data, opt}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRepository)(nil).List), varargs...)
}

// SetMerchantDB mocks base method.
func (m *MockIRepository) SetMerchantDB(ctx context.Context, merchantId uint64, connectStr string, databaseType db.DatabaseType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMerchantDB", ctx, merchantId, connectStr, databaseType)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMerchantDB indicates an expected call of SetMerchantDB.
func (mr *MockIRepositoryMockRecorder) SetMerchantDB(ctx, merchantId, connectStr, databaseType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMerchantDB", reflect.TypeOf((*MockIRepository)(nil).SetMerchantDB), ctx, merchantId, connectStr, databaseType)
}

// Transaction mocks base method.
func (m *MockIRepository) Transaction(ctx context.Context, fc func(*gorm.DB) error, opts ...*sql.TxOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, fc}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockIRepositoryMockRecorder) Transaction(ctx, fc interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, fc}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockIRepository)(nil).Transaction), varargs...)
}

// Update mocks base method.
func (m *MockIRepository) Update(ctx context.Context, tx *gorm.DB, opt iface.WhereOption, col iface.UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, tx, opt, col}
	for _, a := range scopes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIRepositoryMockRecorder) Update(ctx, tx, opt, col interface{}, scopes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, tx, opt, col}, scopes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRepository)(nil).Update), varargs...)
}

// MockIMerchantBaseRepository is a mock of IMerchantBaseRepository interface.
type MockIMerchantBaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMerchantBaseRepositoryMockRecorder
}

// MockIMerchantBaseRepositoryMockRecorder is the mock recorder for MockIMerchantBaseRepository.
type MockIMerchantBaseRepositoryMockRecorder struct {
	mock *MockIMerchantBaseRepository
}

// NewMockIMerchantBaseRepository creates a new mock instance.
func NewMockIMerchantBaseRepository(ctrl *gomock.Controller) *MockIMerchantBaseRepository {
	mock := &MockIMerchantBaseRepository{ctrl: ctrl}
	mock.recorder = &MockIMerchantBaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMerchantBaseRepository) EXPECT() *MockIMerchantBaseRepositoryMockRecorder {
	return m.recorder
}

// MockIMerchantRepository is a mock of IMerchantRepository interface.
type MockIMerchantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMerchantRepositoryMockRecorder
}

// MockIMerchantRepositoryMockRecorder is the mock recorder for MockIMerchantRepository.
type MockIMerchantRepositoryMockRecorder struct {
	mock *MockIMerchantRepository
}

// NewMockIMerchantRepository creates a new mock instance.
func NewMockIMerchantRepository(ctrl *gomock.Controller) *MockIMerchantRepository {
	mock := &MockIMerchantRepository{ctrl: ctrl}
	mock.recorder = &MockIMerchantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMerchantRepository) EXPECT() *MockIMerchantRepositoryMockRecorder {
	return m.recorder
}

// DeleteMerchantDB mocks base method.
func (m *MockIMerchantRepository) DeleteMerchantDB(ctx context.Context, merchantId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMerchantDB", ctx, merchantId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMerchantDB indicates an expected call of DeleteMerchantDB.
func (mr *MockIMerchantRepositoryMockRecorder) DeleteMerchantDB(ctx, merchantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMerchantDB", reflect.TypeOf((*MockIMerchantRepository)(nil).DeleteMerchantDB), ctx, merchantId)
}

// GetALLMerchantDB mocks base method.
func (m *MockIMerchantRepository) GetALLMerchantDB(ctx context.Context) (map[uint64]*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetALLMerchantDB", ctx)
	ret0, _ := ret[0].(map[uint64]*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetALLMerchantDB indicates an expected call of GetALLMerchantDB.
func (mr *MockIMerchantRepositoryMockRecorder) GetALLMerchantDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetALLMerchantDB", reflect.TypeOf((*MockIMerchantRepository)(nil).GetALLMerchantDB), ctx)
}

// GetMerchantDB mocks base method.
func (m *MockIMerchantRepository) GetMerchantDB(ctx context.Context, merchantId uint64) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerchantDB", ctx, merchantId)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMerchantDB indicates an expected call of GetMerchantDB.
func (mr *MockIMerchantRepositoryMockRecorder) GetMerchantDB(ctx, merchantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerchantDB", reflect.TypeOf((*MockIMerchantRepository)(nil).GetMerchantDB), ctx, merchantId)
}

// SetMerchantDB mocks base method.
func (m *MockIMerchantRepository) SetMerchantDB(ctx context.Context, merchantId uint64, connectStr string, databaseType db.DatabaseType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMerchantDB", ctx, merchantId, connectStr, databaseType)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMerchantDB indicates an expected call of SetMerchantDB.
func (mr *MockIMerchantRepositoryMockRecorder) SetMerchantDB(ctx, merchantId, connectStr, databaseType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMerchantDB", reflect.TypeOf((*MockIMerchantRepository)(nil).SetMerchantDB), ctx, merchantId, connectStr, databaseType)
}

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockIUserRepository) GetUserByID(ctx context.Context, userID uint64) (dto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(dto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIUserRepositoryMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByID), ctx, userID)
}

// GetUserIDs mocks base method.
func (m *MockIUserRepository) GetUserIDs(ctx context.Context, opt *option.UserWhereOption) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDs", ctx, opt)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDs indicates an expected call of GetUserIDs.
func (mr *MockIUserRepositoryMockRecorder) GetUserIDs(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDs", reflect.TypeOf((*MockIUserRepository)(nil).GetUserIDs), ctx, opt)
}

// MockICacheRepository is a mock of ICacheRepository interface.
type MockICacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICacheRepositoryMockRecorder
}

// MockICacheRepositoryMockRecorder is the mock recorder for MockICacheRepository.
type MockICacheRepositoryMockRecorder struct {
	mock *MockICacheRepository
}

// NewMockICacheRepository creates a new mock instance.
func NewMockICacheRepository(ctrl *gomock.Controller) *MockICacheRepository {
	mock := &MockICacheRepository{ctrl: ctrl}
	mock.recorder = &MockICacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheRepository) EXPECT() *MockICacheRepositoryMockRecorder {
	return m.recorder
}

// AddUserOnline mocks base method.
func (m *MockICacheRepository) AddUserOnline(ctx context.Context, user *dto.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserOnline", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserOnline indicates an expected call of AddUserOnline.
func (mr *MockICacheRepositoryMockRecorder) AddUserOnline(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserOnline", reflect.TypeOf((*MockICacheRepository)(nil).AddUserOnline), ctx, user)
}

// Del mocks base method.
func (m *MockICacheRepository) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockICacheRepositoryMockRecorder) Del(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockICacheRepository)(nil).Del), ctx, key)
}

// Exists mocks base method.
func (m *MockICacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockICacheRepositoryMockRecorder) Exists(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockICacheRepository)(nil).Exists), ctx, key)
}

// FlushAllCache mocks base method.
func (m *MockICacheRepository) FlushAllCache(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAllCache", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushAllCache indicates an expected call of FlushAllCache.
func (mr *MockICacheRepositoryMockRecorder) FlushAllCache(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAllCache", reflect.TypeOf((*MockICacheRepository)(nil).FlushAllCache), ctx)
}

// Get mocks base method.
func (m *MockICacheRepository) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockICacheRepositoryMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockICacheRepository)(nil).Get), ctx, key)
}

// LLen mocks base method.
func (m *MockICacheRepository) LLen(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LLen", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LLen indicates an expected call of LLen.
func (mr *MockICacheRepositoryMockRecorder) LLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LLen", reflect.TypeOf((*MockICacheRepository)(nil).LLen), ctx, key)
}

// Publish mocks base method.
func (m *MockICacheRepository) Publish(ctx context.Context, key string, message interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, key, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockICacheRepositoryMockRecorder) Publish(ctx, key, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockICacheRepository)(nil).Publish), ctx, key, message)
}

// RPush mocks base method.
func (m *MockICacheRepository) RPush(ctx context.Context, key string, data interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPush", ctx, key, data)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RPush indicates an expected call of RPush.
func (mr *MockICacheRepositoryMockRecorder) RPush(ctx, key, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockICacheRepository)(nil).RPush), ctx, key, data)
}

// RemoveUserOnline mocks base method.
func (m *MockICacheRepository) RemoveUserOnline(ctx context.Context, user *dto.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserOnline", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUserOnline indicates an expected call of RemoveUserOnline.
func (mr *MockICacheRepositoryMockRecorder) RemoveUserOnline(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserOnline", reflect.TypeOf((*MockICacheRepository)(nil).RemoveUserOnline), ctx, user)
}

// Scan mocks base method.
func (m *MockICacheRepository) Scan(ctx context.Context, pattern string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, pattern)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan.
func (mr *MockICacheRepositoryMockRecorder) Scan(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockICacheRepository)(nil).Scan), ctx, pattern)
}

// SetEX mocks base method.
func (m *MockICacheRepository) SetEX(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEX", ctx, key, data, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEX indicates an expected call of SetEX.
func (mr *MockICacheRepositoryMockRecorder) SetEX(ctx, key, data, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEX", reflect.TypeOf((*MockICacheRepository)(nil).SetEX), ctx, key, data, expireAt)
}

// SetEXWithJson mocks base method.
func (m *MockICacheRepository) SetEXWithJson(ctx context.Context, key string, data interface{}, expireAt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEXWithJson", ctx, key, data, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEXWithJson indicates an expected call of SetEXWithJson.
func (mr *MockICacheRepositoryMockRecorder) SetEXWithJson(ctx, key, data, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEXWithJson", reflect.TypeOf((*MockICacheRepository)(nil).SetEXWithJson), ctx, key, data, expireAt)
}

// SetNX mocks base method.
func (m *MockICacheRepository) SetNX(ctx context.Context, key string, data interface{}, expireAt time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, data, expireAt)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetNX indicates an expected call of SetNX.
func (mr *MockICacheRepositoryMockRecorder) SetNX(ctx, key, data, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockICacheRepository)(nil).SetNX), ctx, key, data, expireAt)
}

// SetTTL mocks base method.
func (m *MockICacheRepository) SetTTL(ctx context.Context, key string, expireAt time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTTL", ctx, key, expireAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTTL indicates an expected call of SetTTL.
func (mr *MockICacheRepositoryMockRecorder) SetTTL(ctx, key, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTTL", reflect.TypeOf((*MockICacheRepository)(nil).SetTTL), ctx, key, expireAt)
}

// UserOnlineMap mocks base method.
func (m *MockICacheRepository) UserOnlineMap(ctx context.Context) (map[string]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserOnlineMap", ctx)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserOnlineMap indicates an expected call of UserOnlineMap.
func (mr *MockICacheRepositoryMockRecorder) UserOnlineMap(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserOnlineMap", reflect.TypeOf((*MockICacheRepository)(nil).UserOnlineMap), ctx)
}

// UserOnlineMapWithKey mocks base method.
func (m *MockICacheRepository) UserOnlineMapWithKey(ctx context.Context) (map[string]bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserOnlineMapWithKey", ctx)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserOnlineMapWithKey indicates an expected call of UserOnlineMapWithKey.
func (mr *MockICacheRepositoryMockRecorder) UserOnlineMapWithKey(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserOnlineMapWithKey", reflect.TypeOf((*MockICacheRepository)(nil).UserOnlineMapWithKey), ctx)
}

// ZAddNX mocks base method.
func (m *MockICacheRepository) ZAddNX(ctx context.Context, key string, members ...redis.Z) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNX", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ZAddNX indicates an expected call of ZAddNX.
func (mr *MockICacheRepositoryMockRecorder) ZAddNX(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNX", reflect.TypeOf((*MockICacheRepository)(nil).ZAddNX), varargs...)
}

// ZCard mocks base method.
func (m *MockICacheRepository) ZCard(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZCard indicates an expected call of ZCard.
func (mr *MockICacheRepositoryMockRecorder) ZCard(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockICacheRepository)(nil).ZCard), ctx, key)
}

// ZPopMin mocks base method.
func (m *MockICacheRepository) ZPopMin(ctx context.Context, key string, popCount int64) ([]redis.Z, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZPopMin", ctx, key, popCount)
	ret0, _ := ret[0].([]redis.Z)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZPopMin indicates an expected call of ZPopMin.
func (mr *MockICacheRepositoryMockRecorder) ZPopMin(ctx, key, popCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMin", reflect.TypeOf((*MockICacheRepository)(nil).ZPopMin), ctx, key, popCount)
}

// ZRangeByScore mocks base method.
func (m *MockICacheRepository) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockICacheRepositoryMockRecorder) ZRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockICacheRepository)(nil).ZRangeByScore), ctx, key, opt)
}

// ZRangeWithScore mocks base method.
func (m *MockICacheRepository) ZRangeWithScore(ctx context.Context, key string, start, end int64) ([]redis.Z, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeWithScore", ctx, key, start, end)
	ret0, _ := ret[0].([]redis.Z)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZRangeWithScore indicates an expected call of ZRangeWithScore.
func (mr *MockICacheRepositoryMockRecorder) ZRangeWithScore(ctx, key, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeWithScore", reflect.TypeOf((*MockICacheRepository)(nil).ZRangeWithScore), ctx, key, start, end)
}

// ZRem mocks base method.
func (m *MockICacheRepository) ZRem(ctx context.Context, key string, member interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRem", ctx, key, member)
	ret0, _ := ret[0].(error)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockICacheRepositoryMockRecorder) ZRem(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockICacheRepository)(nil).ZRem), ctx, key, member)
}

// MockWhereOption is a mock of WhereOption interface.
type MockWhereOption struct {
	ctrl     *gomock.Controller
	recorder *MockWhereOptionMockRecorder
}

// MockWhereOptionMockRecorder is the mock recorder for MockWhereOption.
type MockWhereOptionMockRecorder struct {
	mock *MockWhereOption
}

// NewMockWhereOption creates a new mock instance.
func NewMockWhereOption(ctrl *gomock.Controller) *MockWhereOption {
	mock := &MockWhereOption{ctrl: ctrl}
	mock.recorder = &MockWhereOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWhereOption) EXPECT() *MockWhereOptionMockRecorder {
	return m.recorder
}

// IsEmptyWhereOpt mocks base method.
func (m *MockWhereOption) IsEmptyWhereOpt() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmptyWhereOpt")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmptyWhereOpt indicates an expected call of IsEmptyWhereOpt.
func (mr *MockWhereOptionMockRecorder) IsEmptyWhereOpt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmptyWhereOpt", reflect.TypeOf((*MockWhereOption)(nil).IsEmptyWhereOpt))
}

// Page mocks base method.
func (m *MockWhereOption) Page(db *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Page", db)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Page indicates an expected call of Page.
func (mr *MockWhereOptionMockRecorder) Page(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Page", reflect.TypeOf((*MockWhereOption)(nil).Page), db)
}

// Preload mocks base method.
func (m *MockWhereOption) Preload(db *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Preload", db)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Preload indicates an expected call of Preload.
func (mr *MockWhereOptionMockRecorder) Preload(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preload", reflect.TypeOf((*MockWhereOption)(nil).Preload), db)
}

// Sort mocks base method.
func (m *MockWhereOption) Sort(db *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sort", db)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Sort indicates an expected call of Sort.
func (mr *MockWhereOptionMockRecorder) Sort(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockWhereOption)(nil).Sort), db)
}

// TableName mocks base method.
func (m *MockWhereOption) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockWhereOptionMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockWhereOption)(nil).TableName))
}

// Where mocks base method.
func (m *MockWhereOption) Where(db *gorm.DB) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Where", db)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockWhereOptionMockRecorder) Where(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockWhereOption)(nil).Where), db)
}

// WithoutCount mocks base method.
func (m *MockWhereOption) WithoutCount() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithoutCount")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WithoutCount indicates an expected call of WithoutCount.
func (mr *MockWhereOptionMockRecorder) WithoutCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithoutCount", reflect.TypeOf((*MockWhereOption)(nil).WithoutCount))
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// TableName mocks base method.
func (m *MockModel) TableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName.
func (mr *MockModelMockRecorder) TableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableName", reflect.TypeOf((*MockModel)(nil).TableName))
}

// MockUpdateColumns is a mock of UpdateColumns interface.
type MockUpdateColumns struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateColumnsMockRecorder
}

// MockUpdateColumnsMockRecorder is the mock recorder for MockUpdateColumns.
type MockUpdateColumnsMockRecorder struct {
	mock *MockUpdateColumns
}

// NewMockUpdateColumns creates a new mock instance.
func NewMockUpdateColumns(ctrl *gomock.Controller) *MockUpdateColumns {
	mock := &MockUpdateColumns{ctrl: ctrl}
	mock.recorder = &MockUpdateColumnsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateColumns) EXPECT() *MockUpdateColumnsMockRecorder {
	return m.recorder
}

// Columns mocks base method.
func (m *MockUpdateColumns) Columns() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Columns indicates an expected call of Columns.
func (mr *MockUpdateColumnsMockRecorder) Columns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockUpdateColumns)(nil).Columns))
}
