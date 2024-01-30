package controllers_test

import (
	"testing"
	"time"
	"urlshortener/common"
	"urlshortener/controllers"

	"urlshortener/utils"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/jonboulle/clockwork"
)

type controllersTestSuite struct {
	suite.Suite                                    // embed the suite package
	controllers controllers.URLControllerInterface // embed the repo which contains the methods to test
	db          *gorm.DB                           // test database instance
	clock       clockwork.Clock                    // mock clock for freezing time
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
	s.db = common.GetDBTest()
	s.controllers = controllers.NewURLController()
	// Truncate all tables in the database
	if err := utils.TruncateTables(s.db); err != nil {
		s.T().Fatal(err)
	}
}

func TestControllersTestSuite(t *testing.T) {
	suite.Run(t, new(controllersTestSuite))
}
