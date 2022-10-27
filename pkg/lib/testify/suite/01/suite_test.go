package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
	suite.Suite
	testCount uint32
}

func (s *MyTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *MyTestSuite) SetupTest() {
	fmt.Printf("SetupTest test count: %d\n", s.testCount)
}

func (s *MyTestSuite) TearDownTest() {
	s.testCount++
	fmt.Printf("TearDownTest  test count: %d\n", s.testCount)
}

func (s *MyTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest suite:%s test:%s\n", suiteName, testName)
}

func (s *MyTestSuite) TestExample() {
	fmt.Println("TestExample")
}

func TestExample(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
