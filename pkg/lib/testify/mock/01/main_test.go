package _1

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockCrawler struct {
	mock.Mock
}

func (m *MockCrawler) GetUserList() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

var (
	MockUsers []*User
)

func init() {
	MockUsers = append(MockUsers, &User{"dg", 11})
	MockUsers = append(MockUsers, &User{"zs", 22})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil)

	GetAndPrintUsers(crawler)

	crawler.AssertExpectations(t)
}
