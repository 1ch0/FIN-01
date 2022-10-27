package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greeting", greeting)

	mux.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, 200, "get index error")
	assert.Contains(t, recorder.Body.String(), "hello world", "body error")
}

func TestGreeting(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/greeting", nil)
	request.URL.RawQuery = "name=dj"
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greeting", greeting)

	mux.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, 200, "greeting error")
	assert.Contains(t, recorder.Body.String(), "welcome, dj", "body error")
}
