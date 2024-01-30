//=================THIS FILE IS NOT USED, WE WILL BE USING A REAL TEST DB CONNECTION INSTEAD ============================
// experimentation file to see if i could mock a DB

// in this file, we will use gomock to mock the gorm db connection to use for testing
// this is what the gorm package is like and its methods:

// type DB interface {
// 	Model(dest interface{}) DB
// 	Where(query interface{}, args ...interface{}) DB
// 	Count(dest interface{}) DB
// 	Create(dest interface{}) *DB
// 	// Add other methods used by your CRUD repository
// }

package mocks

import (
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

// MockDB
// is a mocked version of the gorm.DB interface
// contains a controller and a recorder
type MockDB struct {
	//embed gorm.DB interface into the mock
	gorm.DB
	// controller manages the mock object
	mockCrtl *gomock.Controller
	// recorder keeps track of what we expect MockDB to do during a test
	// e.g when my code calls the "Where" method, then...this is what should happen
	recorder *MockDBMockRecorder
}

// Recorder struct
// mock: A reference to the MockDB object being recorded
type MockDBMockRecorder struct {
	mock *MockDB
}

// create a new MockDB obj
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{mockCrtl: ctrl}           // associate it with the controller
	mock.recorder = &MockDBMockRecorder{mock} // create a new recorder obj
	return mock
}

// EXPECT returns the recorder of our new MockDB obj to set expectations
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Model mocks the Model method of gorm.DB
func (m *MockDB) Model(dest interface{}) gorm.DB {
	ret := m.mockCrtl.Call(m, "Model", dest)
	return ret[0].(gorm.DB)
}

// Where mocks the Where method of gorm.DB
func (m *MockDB) Where(query interface{}, args ...interface{}) gorm.DB {
	ret := m.mockCrtl.Call(m, "Where", query, args)
	return ret[0].(gorm.DB)
}

// Count mocks the Count method of gorm.DB
func (m *MockDB) Count(dest interface{}) gorm.DB {
	ret := m.mockCrtl.Call(m, "Count", dest)
	return ret[0].(gorm.DB)
}

// Create mocks the Create method of gorm.DB
func (m *MockDB) Create(dest interface{}) *gorm.DB {
	ret := m.mockCrtl.Call(m, "Create", dest)
	return ret[0].(*gorm.DB)
}

// this here is for the crud_test.go but we never use

// package urlrepo_test

// import (
// 	"testing"
// 	"urlshortener/models/urlmodel"
// 	"urlshortener/common/mocks"
// 	"urlshortener/repo/urlrepo"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/gorm"
// )

// // Create the Test Suite for my crud.go file
// type crudRepositoryTestSuite struct {
// 	suite.Suite                                 // embed the suite package
// 	crudRepo        urlrepo.CRUDRepositoryInterface // methods to test
// 	mockCtrl   *gomock.Controller // controller creates and manage mocks, specify expectations and assert calls
// 	mockDB     *mocks.MockDB // to add our mock DB connection for testing
// }

// // Setup the Test Suite
// func (suite *crudRepositoryTestSuite) SetupSuite() {
// 	suite.mockCtrl = gomock.NewController(suite.T())
// 	suite.mockDB = mocks.NewMockDB(suite.mockCtrl)
// 	suite.crudRepo = urlrepo.NewCRUDRepository(suite.mockDB)
// }

// // Tear down the Test Suite
// func (suite *crudRepositoryTestSuite) TearDownSuite() {
// 	suite.mockCtrl.Finish()
// }

// // Test cases go here
// // Test case 1: Invalid DB connection
// func (suite *crudRepositoryTestSuite) TestInvalidDBConnection() {
// 	suite.mockDB.EXPECT().Create(gomock.Any()).Return(gorm.ErrInvalidDB)
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{})
// 	suite.Error(err)
// }

// // Test case 2: Invalid URL model
// func (suite *crudRepositoryTestSuite) TestInvalidURLModel() {
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{})
// 	suite.Error(err)
// }

// // Test case 3: Valid result returned
// func (suite *crudRepositoryTestSuite) TestValidResultReturned() {
// 	suite.mockDB.EXPECT().Create(gomock.Any()).Return(nil)
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{ShortURL: "existingShortURL"})
// 	suite.NoError(err)
// }

// // Test case 4: Invalid result returned
// func (suite *crudRepositoryTestSuite) TestInvalidResultReturned() {
// 	suite.mockDB.EXPECT().Create(gomock.Any()).Return(fmt.Errorf("database error"))
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{})
// 	suite.Error(err)
// }

// // Test case 5: Count is not > 0
// func (suite *crudRepositoryTestSuite) TestCountNotGreaterThanZero() {
// 	gomock.InOrder(
// 		suite.mockDB.EXPECT().Model(gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Count(gomock.Any()).Return(suite.mockDB, int64(0)),
// 	)
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{ShortURL: "nonExistingShortURL"})
// 	suite.NoError(err)
// }

// // Test case 6: Insert record fails
// func (suite *crudRepositoryTestSuite) TestInsertRecordFails() {
// 	gomock.InOrder(
// 		suite.mockDB.EXPECT().Model(gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Count(gomock.Any()).Return(suite.mockDB, int64(0)),
// 		suite.mockDB.EXPECT().Create(gomock.Any()).Return(fmt.Errorf("database error")),
// 	)
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{ShortURL: "nonExistingShortURL"})
// 	suite.Error(err)
// }

// // Test case 7: Valid test case, function goes through
// func (suite *crudRepositoryTestSuite) TestValidTestCase() {
// 	gomock.InOrder(
// 		suite.mockDB.EXPECT().Model(gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Where(gomock.Any(), gomock.Any()).Return(suite.mockDB),
// 		suite.mockDB.EXPECT().Count(gomock.Any()).Return(suite.mockDB, int64(0)),
// 		suite.mockDB.EXPECT().Create(gomock.Any()).Return(nil),
// 	)
// 	err := suite.crudRepo.StoreURL(urlmodel.URL{ShortURL: "nonExistingShortURL"})
// 	suite.NoError(err)
// }

// // Run the Test Suite
// func TestCRUDRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t, new(crudRepositoryTestSuite))
// }
