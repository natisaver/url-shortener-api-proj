package urlrepo_test

import (
	"testing"
	"time"
	"urlshortener/common"
	"urlshortener/models/urlmodel"
	"urlshortener/repo/urlrepo"
	"urlshortener/utils"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/jonboulle/clockwork"
)

// Create the Test Suite for my crud.go file
// https://pkg.go.dev/testing

type crudRepositoryTestSuite struct {
	suite.Suite                                 // embed the suite package
	crudRepo    urlrepo.CRUDRepositoryInterface // embed the repo which contains the methods to test
	db          *gorm.DB                        // test database instance
	clock       clockwork.Clock                 // mock clock for freezing time

}

// Setup the Test Suite
func (s *crudRepositoryTestSuite) SetupSuite() {
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
func (s *crudRepositoryTestSuite) SetupTest() {
	// Create CRUD repository instance using the test database
	s.db = common.GetDBTest()
	s.crudRepo = urlrepo.NewCRUDRepository(s.db)
	// Truncate all tables in the database
	if err := utils.TruncateTables(s.db); err != nil {
		s.T().Fatal(err)
	}
}

// if required:
// TearDownSuite()
// TearDownTest()

// Run the Test Suite
// testing.T stores the state of the test.
// You call methods on it like t.Fail to say the test failed, or t.Skip to say the test was skipped, etc. It remembers all this, and Go uses it to report what happened in all the test functions.
func TestCRUDRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(crudRepositoryTestSuite))
}

// Test cases go here
// https://gorm.io/docs/create.html
// https://pkg.go.dev/github.com/stretchr/testify/assert#GreaterOrEqual

// StoreURL
func (s *crudRepositoryTestSuite) TestStoreURLModelNotExists() {
	err := s.crudRepo.StoreURL(urlmodel.URL{})
	assert.ErrorContains(s.T(), err, "empty object")
}

func (s *crudRepositoryTestSuite) TestStoreURLModelShortURLDuplicate() {
	var count int64
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	mockURLNew := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcde",
	}

	_ = s.db.Create(&mockURL)
	err := s.crudRepo.StoreURL(mockURLNew)
	_ = s.db.Model(&urlmodel.URL{}).Where("shorturl = ?", mockURL.ShortURL).Count(&count)

	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), count, 1)
}

func (s *crudRepositoryTestSuite) TestStoreURLInsert() {
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	//insert
	_ = s.crudRepo.StoreURL(mockURL)
	//get
	res := s.db.First("urls")
	//not nil
	assert.NotNil(s.T(), res)
}

// GetURL
func (s *crudRepositoryTestSuite) TestGetURLModelNotExists() {
	_, err := s.crudRepo.GetURL(urlmodel.URL{})
	assert.ErrorContains(s.T(), err, "empty object")
}

func (s *crudRepositoryTestSuite) TestGetURLModelEmptyDB() {
	// shorturl to search not in db
	mockURLNew := urlmodel.URL{
		ShortURL: "abcde",
	}

	//get
	res, err := s.crudRepo.GetURL(mockURLNew)
	assert.EqualValues(s.T(), err, nil)
	assert.EqualValues(s.T(), res, "")

}

func (s *crudRepositoryTestSuite) TestGetURLModelRecordNotExists() {
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcd",
	}
	// shorturl to search not in db
	mockURLNew := urlmodel.URL{
		ShortURL: "abcde",
	}

	_ = s.db.Create(&mockURL)
	//get
	res, err := s.crudRepo.GetURL(mockURLNew)
	assert.EqualValues(s.T(), err, nil)
	assert.EqualValues(s.T(), res, "")

}

func (s *crudRepositoryTestSuite) TestGetURLModelLongURLNotExists() {
	mockURL := urlmodel.URL{
		ShortURL: "abc",
	}
	// longurl to get not in db
	mockURLNew := urlmodel.URL{
		ShortURL: "abc",
	}

	_ = s.db.Create(&mockURL)
	//get
	res, err := s.crudRepo.GetURL(mockURLNew)
	assert.EqualValues(s.T(), err, nil)
	assert.EqualValues(s.T(), res, "")
}

func (s *crudRepositoryTestSuite) TestGetURLInsert() {
	mockURL := urlmodel.URL{
		ShortURL: "abc",
		LongURL:  "abcde",
	}
	//insert
	_ = s.db.Create(&mockURL)
	//get
	res, err := s.crudRepo.GetURL(mockURL)
	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), res, "abcde")
}
