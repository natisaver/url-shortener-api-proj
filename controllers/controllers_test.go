package controllers_test

import (
	"context"
	"errors"
	"testing"
	"time"
	"urlshortener/common"
	"urlshortener/config"
	"urlshortener/controllers"
	"urlshortener/models/urlmodel"
	mock_urlrepo "urlshortener/repo/urlrepo/mocks"

	"urlshortener/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/jonboulle/clockwork"
)

type controllersTestSuite struct {
	suite.Suite                                    // embed the suite package
	controllers controllers.URLControllerInterface // embed the repo which contains the methods to test
	db          *gorm.DB                           // test database instance
	clock       clockwork.Clock                    // mock clock for freezing time
	// Generally, it is not good to put ctx in struct as that reduces flexibility in changing it:
	// e.g if I want to pass 2 different ctx into funcA and funcB, i will have to instantiate 2 different objects
	// Instead, define ctx in the function input params e.g., funcA(ctxA) and funcB(ctxB)
	// Now, I can now have more flexibility to change the ctx before i pass it into functions
	// However, for testing in a suite, it makes sense to put in a struct since i want a baseline context before each test
	ctx context.Context
}

// Setup the Test Suite
func (s *controllersTestSuite) SetupSuite() {
	// init test db
	_, err := common.InitDBTest()
	if err != nil {
		panic(err)
	}
	// Set the fake clock to a specific time
	// time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	frozenTime := time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC)

	// init mock clock
	s.clock = clockwork.NewFakeClockAt(frozenTime)
}

// Run before each test
func (s *controllersTestSuite) SetupTest() {
	// set up fresh context before each test
	s.ctx = context.Background()
	s.db = common.GetDBTest()
	// and add the test db into the context
	s.ctx = context.WithValue(s.ctx, config.CtxKeyDB, s.db)
	s.controllers = controllers.NewURLController()
	// Truncate all tables in the database
	if err := utils.TruncateTables(s.db); err != nil {
		s.T().Fatal(err)
	}
}

func TestControllersTestSuite(t *testing.T) {
	suite.Run(t, new(controllersTestSuite))
}

// Test cases go here

// StoreURLController
func (s *controllersTestSuite) Test_StoreURLController_CRUDFail() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	mockURL := urlmodel.URL{}

	// set DB to testDB
	// option 1, change variable function getDB's implementation to the test DB
	// common.GetDB = common.GetDBTest

	// option 2, take test db from context and pass into controller methods
	ctx := s.ctx

	// pass the mock StoreURL to context
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	mockCRUDRepository.EXPECT().StoreURL(gomock.Any()).Return(errors.New("empty object"))
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)

	// function under test
	err := s.controllers.StoreURLController(ctx, mockURL)

	assert.EqualValues(s.T(), err, errors.New("empty object"))
}

func (s *controllersTestSuite) Test_StoreURLController_Panic() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	mockURL := urlmodel.URL{}
	ctx := s.ctx
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	// The Do method expects a function with the same signature as the method you are mocking
	// in this case, we force a panic to occur
	mockCRUDRepository.EXPECT().StoreURL(gomock.Any()).Do(func(_ urlmodel.URL) {
		panic("exception")
	})
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)
	// then we see if it panics
	// this assert.Panics works even if you have StopPanic in the function you are testing
	// as long as a panic did occur it will test for that
	assert.Panics(s.T(), func() { s.controllers.StoreURLController(ctx, mockURL) }, "Function did not Panic")
}

func (s *controllersTestSuite) Test_StoreURLController_CRUDSuccess() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	ctx := s.ctx
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	// didnt add param .AnyTimes, as we expect it to run once
	mockCRUDRepository.EXPECT().StoreURL(gomock.Eq(mockURL)).Return(nil)
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)
	err := s.controllers.StoreURLController(ctx, mockURL)
	assert.Nil(s.T(), err)
}

// GetLongURLController

func (s *controllersTestSuite) Test_GetLongURLController_Panic() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	ctx := s.ctx
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	mockCRUDRepository.EXPECT().GetURL(gomock.Any()).AnyTimes().Do(func(_ urlmodel.URL) {
		panic("exception")
	})
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)
	assert.Panics(s.T(), func() { s.controllers.GetLongURLController(ctx, mockURL) }, "Function did not Panic")
}

func (s *controllersTestSuite) Test_GetLongURLController_CRUDFailure() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	ctx := s.ctx
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	mockCRUDRepository.EXPECT().GetURL(gomock.Eq(mockURL)).AnyTimes().Return("", errors.New("Not found"))
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)
	longurl, err := s.controllers.GetLongURLController(ctx, mockURL)
	assert.NotNil(s.T(), err)
	assert.EqualValues(s.T(), longurl, "")
}

func (s *controllersTestSuite) Test_GetLongURLController_CRUDSuccess() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	ctx := s.ctx
	mockCRUDRepository := mock_urlrepo.NewMockCRUDRepositoryInterface(mockCtrl)
	mockCRUDRepository.EXPECT().GetURL(gomock.Eq(mockURL)).Return("abcd", nil)
	ctx = context.WithValue(ctx, config.CtxKeyMockCRUDRepository, mockCRUDRepository)
	longurl, err := s.controllers.GetLongURLController(ctx, mockURL)
	assert.EqualValues(s.T(), longurl, "abcd")
	assert.Nil(s.T(), err)
}
