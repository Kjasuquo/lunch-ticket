// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	multipart "mime/multipart"
	reflect "reflect"
	time "time"

	session "github.com/aws/aws-sdk-go/aws/session"
	models "github.com/decadevs/lunch-api/internal/core/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddTokenToBlacklist mocks base method.
func (m *MockUserRepository) AddTokenToBlacklist(email, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTokenToBlacklist", email, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTokenToBlacklist indicates an expected call of AddTokenToBlacklist.
func (mr *MockUserRepositoryMockRecorder) AddTokenToBlacklist(email, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTokenToBlacklist", reflect.TypeOf((*MockUserRepository)(nil).AddTokenToBlacklist), email, token)
}

// AdminEmailVerification mocks base method.
func (m *MockUserRepository) AdminEmailVerification(id string) (*models.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminEmailVerification", id)
	ret0, _ := ret[0].(*models.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminEmailVerification indicates an expected call of AdminEmailVerification.
func (mr *MockUserRepositoryMockRecorder) AdminEmailVerification(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminEmailVerification", reflect.TypeOf((*MockUserRepository)(nil).AdminEmailVerification), id)
}

// AdminResetPassword mocks base method.
func (m *MockUserRepository) AdminResetPassword(id, newPassword string) (*models.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminResetPassword", id, newPassword)
	ret0, _ := ret[0].(*models.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminResetPassword indicates an expected call of AdminResetPassword.
func (mr *MockUserRepositoryMockRecorder) AdminResetPassword(id, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminResetPassword", reflect.TypeOf((*MockUserRepository)(nil).AdminResetPassword), id, newPassword)
}

// CreateAdmin mocks base method.
func (m *MockUserRepository) CreateAdmin(user *models.Admin) (*models.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", user)
	ret0, _ := ret[0].(*models.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockUserRepositoryMockRecorder) CreateAdmin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockUserRepository)(nil).CreateAdmin), user)
}

// CreateFoodBenefactor mocks base method.
func (m *MockUserRepository) CreateFoodBenefactor(user *models.FoodBeneficiary) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFoodBenefactor", user)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFoodBenefactor indicates an expected call of CreateFoodBenefactor.
func (mr *MockUserRepositoryMockRecorder) CreateFoodBenefactor(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFoodBenefactor", reflect.TypeOf((*MockUserRepository)(nil).CreateFoodBenefactor), user)
}

// CreateFoodBenefactorBrunchMealRecord mocks base method.
func (m *MockUserRepository) CreateFoodBenefactorBrunchMealRecord(user *models.FoodBeneficiary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFoodBenefactorBrunchMealRecord", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFoodBenefactorBrunchMealRecord indicates an expected call of CreateFoodBenefactorBrunchMealRecord.
func (mr *MockUserRepositoryMockRecorder) CreateFoodBenefactorBrunchMealRecord(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFoodBenefactorBrunchMealRecord", reflect.TypeOf((*MockUserRepository)(nil).CreateFoodBenefactorBrunchMealRecord), user)
}

// CreateFoodBenefactorDinnerMealRecord mocks base method.
func (m *MockUserRepository) CreateFoodBenefactorDinnerMealRecord(user *models.FoodBeneficiary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFoodBenefactorDinnerMealRecord", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFoodBenefactorDinnerMealRecord indicates an expected call of CreateFoodBenefactorDinnerMealRecord.
func (mr *MockUserRepositoryMockRecorder) CreateFoodBenefactorDinnerMealRecord(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFoodBenefactorDinnerMealRecord", reflect.TypeOf((*MockUserRepository)(nil).CreateFoodBenefactorDinnerMealRecord), user)
}

// CreateFoodTimetable mocks base method.
func (m *MockUserRepository) CreateFoodTimetable(food models.Food) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFoodTimetable", food)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFoodTimetable indicates an expected call of CreateFoodTimetable.
func (mr *MockUserRepositoryMockRecorder) CreateFoodTimetable(food interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFoodTimetable", reflect.TypeOf((*MockUserRepository)(nil).CreateFoodTimetable), food)
}

// CreateKitchenStaff mocks base method.
func (m *MockUserRepository) CreateKitchenStaff(user *models.KitchenStaff) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKitchenStaff", user)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKitchenStaff indicates an expected call of CreateKitchenStaff.
func (mr *MockUserRepositoryMockRecorder) CreateKitchenStaff(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKitchenStaff", reflect.TypeOf((*MockUserRepository)(nil).CreateKitchenStaff), user)
}

// DeleteMeal mocks base method.
func (m *MockUserRepository) DeleteMeal(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMeal", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMeal indicates an expected call of DeleteMeal.
func (mr *MockUserRepositoryMockRecorder) DeleteMeal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMeal", reflect.TypeOf((*MockUserRepository)(nil).DeleteMeal), id)
}

// FindAdminByEmail mocks base method.
func (m *MockUserRepository) FindAdminByEmail(email string) (*models.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAdminByEmail", email)
	ret0, _ := ret[0].(*models.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAdminByEmail indicates an expected call of FindAdminByEmail.
func (mr *MockUserRepositoryMockRecorder) FindAdminByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAdminByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindAdminByEmail), email)
}

// FindAllFoodBeneficiary mocks base method.
func (m *MockUserRepository) FindAllFoodBeneficiary() ([]models.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFoodBeneficiary")
	ret0, _ := ret[0].([]models.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFoodBeneficiary indicates an expected call of FindAllFoodBeneficiary.
func (mr *MockUserRepositoryMockRecorder) FindAllFoodBeneficiary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFoodBeneficiary", reflect.TypeOf((*MockUserRepository)(nil).FindAllFoodBeneficiary))
}

// FindBrunchByDate mocks base method.
func (m *MockUserRepository) FindBrunchByDate(year int, month time.Month, day int) ([]models.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBrunchByDate", year, month, day)
	ret0, _ := ret[0].([]models.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBrunchByDate indicates an expected call of FindBrunchByDate.
func (mr *MockUserRepositoryMockRecorder) FindBrunchByDate(year, month, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBrunchByDate", reflect.TypeOf((*MockUserRepository)(nil).FindBrunchByDate), year, month, day)
}

// FindDinnerByDate mocks base method.
func (m *MockUserRepository) FindDinnerByDate(year int, month time.Month, day int) ([]models.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDinnerByDate", year, month, day)
	ret0, _ := ret[0].([]models.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindDinnerByDate indicates an expected call of FindDinnerByDate.
func (mr *MockUserRepositoryMockRecorder) FindDinnerByDate(year, month, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDinnerByDate", reflect.TypeOf((*MockUserRepository)(nil).FindDinnerByDate), year, month, day)
}

// FindFoodBenefactorByEmail mocks base method.
func (m *MockUserRepository) FindFoodBenefactorByEmail(email string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFoodBenefactorByEmail", email)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFoodBenefactorByEmail indicates an expected call of FindFoodBenefactorByEmail.
func (mr *MockUserRepositoryMockRecorder) FindFoodBenefactorByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFoodBenefactorByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindFoodBenefactorByEmail), email)
}

// FindFoodBenefactorByFullName mocks base method.
func (m *MockUserRepository) FindFoodBenefactorByFullName(fullname string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFoodBenefactorByFullName", fullname)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFoodBenefactorByFullName indicates an expected call of FindFoodBenefactorByFullName.
func (mr *MockUserRepositoryMockRecorder) FindFoodBenefactorByFullName(fullname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFoodBenefactorByFullName", reflect.TypeOf((*MockUserRepository)(nil).FindFoodBenefactorByFullName), fullname)
}

// FindFoodBenefactorByLocation mocks base method.
func (m *MockUserRepository) FindFoodBenefactorByLocation(location string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFoodBenefactorByLocation", location)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFoodBenefactorByLocation indicates an expected call of FindFoodBenefactorByLocation.
func (mr *MockUserRepositoryMockRecorder) FindFoodBenefactorByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFoodBenefactorByLocation", reflect.TypeOf((*MockUserRepository)(nil).FindFoodBenefactorByLocation), location)
}

// FindFoodBenefactorMealRecord mocks base method.
func (m *MockUserRepository) FindFoodBenefactorMealRecord(email, date string) (*models.MealRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFoodBenefactorMealRecord", email, date)
	ret0, _ := ret[0].(*models.MealRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFoodBenefactorMealRecord indicates an expected call of FindFoodBenefactorMealRecord.
func (mr *MockUserRepositoryMockRecorder) FindFoodBenefactorMealRecord(email, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFoodBenefactorMealRecord", reflect.TypeOf((*MockUserRepository)(nil).FindFoodBenefactorMealRecord), email, date)
}

// FindKitchenStaffByEmail mocks base method.
func (m *MockUserRepository) FindKitchenStaffByEmail(email string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKitchenStaffByEmail", email)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKitchenStaffByEmail indicates an expected call of FindKitchenStaffByEmail.
func (mr *MockUserRepositoryMockRecorder) FindKitchenStaffByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKitchenStaffByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindKitchenStaffByEmail), email)
}

// FindKitchenStaffByFullName mocks base method.
func (m *MockUserRepository) FindKitchenStaffByFullName(fullname string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKitchenStaffByFullName", fullname)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKitchenStaffByFullName indicates an expected call of FindKitchenStaffByFullName.
func (mr *MockUserRepositoryMockRecorder) FindKitchenStaffByFullName(fullname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKitchenStaffByFullName", reflect.TypeOf((*MockUserRepository)(nil).FindKitchenStaffByFullName), fullname)
}

// FindKitchenStaffByLocation mocks base method.
func (m *MockUserRepository) FindKitchenStaffByLocation(location string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKitchenStaffByLocation", location)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKitchenStaffByLocation indicates an expected call of FindKitchenStaffByLocation.
func (mr *MockUserRepositoryMockRecorder) FindKitchenStaffByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKitchenStaffByLocation", reflect.TypeOf((*MockUserRepository)(nil).FindKitchenStaffByLocation), location)
}

// FindUserById mocks base method.
func (m *MockUserRepository) FindUserById(id string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", id)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserRepositoryMockRecorder) FindUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserRepository)(nil).FindUserById), id)
}

// FoodBeneficiaryEmailVerification mocks base method.
func (m *MockUserRepository) FoodBeneficiaryEmailVerification(id string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FoodBeneficiaryEmailVerification", id)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FoodBeneficiaryEmailVerification indicates an expected call of FoodBeneficiaryEmailVerification.
func (mr *MockUserRepositoryMockRecorder) FoodBeneficiaryEmailVerification(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FoodBeneficiaryEmailVerification", reflect.TypeOf((*MockUserRepository)(nil).FoodBeneficiaryEmailVerification), id)
}

// GetFoodByID mocks base method.
func (m *MockUserRepository) GetFoodByID(id string) (*models.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoodByID", id)
	ret0, _ := ret[0].(*models.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoodByID indicates an expected call of GetFoodByID.
func (mr *MockUserRepositoryMockRecorder) GetFoodByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoodByID", reflect.TypeOf((*MockUserRepository)(nil).GetFoodByID), id)
}

// GetTotalUsers mocks base method.
func (m *MockUserRepository) GetTotalUsers() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalUsers")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalUsers indicates an expected call of GetTotalUsers.
func (mr *MockUserRepositoryMockRecorder) GetTotalUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalUsers", reflect.TypeOf((*MockUserRepository)(nil).GetTotalUsers))
}

// KitchenStaffEmailVerification mocks base method.
func (m *MockUserRepository) KitchenStaffEmailVerification(id string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KitchenStaffEmailVerification", id)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KitchenStaffEmailVerification indicates an expected call of KitchenStaffEmailVerification.
func (mr *MockUserRepositoryMockRecorder) KitchenStaffEmailVerification(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KitchenStaffEmailVerification", reflect.TypeOf((*MockUserRepository)(nil).KitchenStaffEmailVerification), id)
}

// KitchenStaffResetPassword mocks base method.
func (m *MockUserRepository) KitchenStaffResetPassword(id, newPassword string) (*models.KitchenStaff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KitchenStaffResetPassword", id, newPassword)
	ret0, _ := ret[0].(*models.KitchenStaff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KitchenStaffResetPassword indicates an expected call of KitchenStaffResetPassword.
func (mr *MockUserRepositoryMockRecorder) KitchenStaffResetPassword(id, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KitchenStaffResetPassword", reflect.TypeOf((*MockUserRepository)(nil).KitchenStaffResetPassword), id, newPassword)
}

// SearchFoodBeneficiary mocks base method.
func (m *MockUserRepository) SearchFoodBeneficiary(text string) ([]models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFoodBeneficiary", text)
	ret0, _ := ret[0].([]models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFoodBeneficiary indicates an expected call of SearchFoodBeneficiary.
func (mr *MockUserRepositoryMockRecorder) SearchFoodBeneficiary(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFoodBeneficiary", reflect.TypeOf((*MockUserRepository)(nil).SearchFoodBeneficiary), text)
}

// TokenInBlacklist mocks base method.
func (m *MockUserRepository) TokenInBlacklist(token *string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenInBlacklist", token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// TokenInBlacklist indicates an expected call of TokenInBlacklist.
func (mr *MockUserRepositoryMockRecorder) TokenInBlacklist(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenInBlacklist", reflect.TypeOf((*MockUserRepository)(nil).TokenInBlacklist), token)
}

// UpdateFoodBenefactorBrunchMealRecord mocks base method.
func (m *MockUserRepository) UpdateFoodBenefactorBrunchMealRecord(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFoodBenefactorBrunchMealRecord", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFoodBenefactorBrunchMealRecord indicates an expected call of UpdateFoodBenefactorBrunchMealRecord.
func (mr *MockUserRepositoryMockRecorder) UpdateFoodBenefactorBrunchMealRecord(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFoodBenefactorBrunchMealRecord", reflect.TypeOf((*MockUserRepository)(nil).UpdateFoodBenefactorBrunchMealRecord), email)
}

// UpdateFoodBenefactorDinnerMealRecord mocks base method.
func (m *MockUserRepository) UpdateFoodBenefactorDinnerMealRecord(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFoodBenefactorDinnerMealRecord", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFoodBenefactorDinnerMealRecord indicates an expected call of UpdateFoodBenefactorDinnerMealRecord.
func (mr *MockUserRepositoryMockRecorder) UpdateFoodBenefactorDinnerMealRecord(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFoodBenefactorDinnerMealRecord", reflect.TypeOf((*MockUserRepository)(nil).UpdateFoodBenefactorDinnerMealRecord), email)
}

// UpdateMeal mocks base method.
func (m *MockUserRepository) UpdateMeal(id string, food models.Food) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeal", id, food)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeal indicates an expected call of UpdateMeal.
func (mr *MockUserRepositoryMockRecorder) UpdateMeal(id, food interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeal", reflect.TypeOf((*MockUserRepository)(nil).UpdateMeal), id, food)
}

// UpdateStatus mocks base method.
func (m *MockUserRepository) UpdateStatus(food []models.Food, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", food, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockUserRepositoryMockRecorder) UpdateStatus(food, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockUserRepository)(nil).UpdateStatus), food, status)
}

// UserResetPassword mocks base method.
func (m *MockUserRepository) UserResetPassword(id, newPassword string) (*models.FoodBeneficiary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserResetPassword", id, newPassword)
	ret0, _ := ret[0].(*models.FoodBeneficiary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserResetPassword indicates an expected call of UserResetPassword.
func (mr *MockUserRepositoryMockRecorder) UserResetPassword(id, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserResetPassword", reflect.TypeOf((*MockUserRepository)(nil).UserResetPassword), id, newPassword)
}

// MockMailerRepository is a mock of MailerRepository interface.
type MockMailerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMailerRepositoryMockRecorder
}

// MockMailerRepositoryMockRecorder is the mock recorder for MockMailerRepository.
type MockMailerRepositoryMockRecorder struct {
	mock *MockMailerRepository
}

// NewMockMailerRepository creates a new mock instance.
func NewMockMailerRepository(ctrl *gomock.Controller) *MockMailerRepository {
	mock := &MockMailerRepository{ctrl: ctrl}
	mock.recorder = &MockMailerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailerRepository) EXPECT() *MockMailerRepositoryMockRecorder {
	return m.recorder
}

// DecodeToken mocks base method.
func (m *MockMailerRepository) DecodeToken(token, secret string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeToken", token, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeToken indicates an expected call of DecodeToken.
func (mr *MockMailerRepositoryMockRecorder) DecodeToken(token, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeToken", reflect.TypeOf((*MockMailerRepository)(nil).DecodeToken), token, secret)
}

// GenerateNonAuthToken mocks base method.
func (m *MockMailerRepository) GenerateNonAuthToken(UserEmail, secret string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateNonAuthToken", UserEmail, secret)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateNonAuthToken indicates an expected call of GenerateNonAuthToken.
func (mr *MockMailerRepositoryMockRecorder) GenerateNonAuthToken(UserEmail, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateNonAuthToken", reflect.TypeOf((*MockMailerRepository)(nil).GenerateNonAuthToken), UserEmail, secret)
}

// SendMail mocks base method.
func (m *MockMailerRepository) SendMail(subject, body, to, Private, Domain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMail", subject, body, to, Private, Domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMail indicates an expected call of SendMail.
func (mr *MockMailerRepositoryMockRecorder) SendMail(subject, body, to, Private, Domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMail", reflect.TypeOf((*MockMailerRepository)(nil).SendMail), subject, body, to, Private, Domain)
}

// MockAWSRepository is a mock of AWSRepository interface.
type MockAWSRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAWSRepositoryMockRecorder
}

// MockAWSRepositoryMockRecorder is the mock recorder for MockAWSRepository.
type MockAWSRepositoryMockRecorder struct {
	mock *MockAWSRepository
}

// NewMockAWSRepository creates a new mock instance.
func NewMockAWSRepository(ctrl *gomock.Controller) *MockAWSRepository {
	mock := &MockAWSRepository{ctrl: ctrl}
	mock.recorder = &MockAWSRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAWSRepository) EXPECT() *MockAWSRepositoryMockRecorder {
	return m.recorder
}

// UploadFileToS3 mocks base method.
func (m *MockAWSRepository) UploadFileToS3(h *session.Session, file multipart.File, fileName string, size int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFileToS3", h, file, fileName, size)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFileToS3 indicates an expected call of UploadFileToS3.
func (mr *MockAWSRepositoryMockRecorder) UploadFileToS3(h, file, fileName, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFileToS3", reflect.TypeOf((*MockAWSRepository)(nil).UploadFileToS3), h, file, fileName, size)
}
