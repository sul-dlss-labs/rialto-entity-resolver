package handlers

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

// MockedRepo is a mocked object that implements the Repository interface
type MockedRepo struct {
	mock.Mock
}

func (m *MockedRepo) QueryForPersonBySunetid(sunetid string) (*string, error) {
	args := m.Called(sunetid)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForPersonByOrcid(orcid string) (*string, error) {
	args := m.Called(orcid)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForPersonByName(firstName string, lastName string) (*string, error) {
	args := m.Called(firstName, lastName)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForOrganizationByName(name string) (*string, error) {
	args := m.Called(name)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForTopicByName(name string) (*string, error) {
	args := m.Called(name)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockedRepo) QueryForGrantByIdentifier(identifier string) (*string, error) {
	args := m.Called(identifier)
	result := args.Get(0)
	if result != nil {
		return result.(*string), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestLookupUserByName(t *testing.T) {
	os.Setenv("API_KEY", "abcdefg")
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForPersonByName", "Aaron", "Collier").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestLookupUserBySunetid(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForPersonBySunetid", "amcollie").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&sunetid=amcollie").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestLookupUserByOrcid(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	id := "http://sul.stanford.edu/rialto/agents/people/8de0ce5e-a2a4-4e61-974e-df6c213cf220"
	repo.On("QueryForPersonByOrcid", "0000-0000-0000-0012").
		Return(&id, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&orcid=0000-0000-0000-0012").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
			})
}

func TestUserNotFound(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForPersonBySunetid", "amcollie").
		Return(nil, nil)
	repo.On("QueryForPersonByOrcid", "0000-0000-0000-0012").
		Return(nil, nil)
	repo.On("QueryForPersonByName", "Aaron", "Collier").
		Return(nil, nil)
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&orcid=0000-0000-0000-0012").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusNotFound, r.Code)
			})
}

func TestUserServerError(t *testing.T) {
	r := gofight.New()
	repo := new(MockedRepo)
	repo.On("QueryForPersonByOrcid", "0000-0000-0000-0012").
		Return(nil, nil)
	repo.On("QueryForPersonByName", "Aaron", "Collier").
		Return(nil, errors.New("Ooops"))
	registry := &runtime.Registry{Repository: repo}
	r.GET("/person?last_name=Collier&first_name=Aaron&orcid=0000-0000-0000-0012").
		SetHeader(gofight.H{
			"X-API-Key": "abcdefg",
		}).
		Run(BuildAPI(registry).Serve(nil),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)
			})
}
