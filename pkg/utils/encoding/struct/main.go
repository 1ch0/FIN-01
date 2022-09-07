package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
)

type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

func main() {
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	m := structs.Map(server)
	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	fmt.Println(m)
}
