package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/suite"
)

type MySuite struct {
	suite.Suite
	recorder *httptest.ResponseRecorder
	mux      *http.ServeMux
}

func (s *MySuite) SetupSuite() {
	s.recorder = httptest.NewRecorder()
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", index)
	s.mux.HandleFunc("/greeting", greeting)
}

func (s *MySuite) TestIndex() {
	request, _ := http.NewRequest("GET", "/", nil)
	s.mux.ServeHTTP(s.recorder, request)

	s.Assert().Equal(s.recorder.Code, 200, "get index error")
	s.Assert().Contains(s.recorder.Body.String(), "Hello World", "body error")
}

func (s *MySuite) TestGreeting() {
	request, _ := http.NewRequest("GET", "/greeting", nil)
	request.URL.RawQuery = "name=dj"

	s.mux.ServeHTTP(s.recorder, request)

	s.Assert().Equal(s.recorder.Code, 200, "greeting error")
	s.Assert().Contains(s.recorder.Body.String(), "welcome, dj", "body error")
}
