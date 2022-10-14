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

	// Create a new struct type:
	s := structs.New(server)

	m := s.Map() // Get a map[string]interface{}
	fmt.Println(m)
	v := s.Values() // Get a []interface{}
	fmt.Println(v)
	f := s.Fields() // Get a []*Field
	fmt.Println(f)
	n := s.Names() // Get a []string
	fmt.Println(n)
	f1 := s.Field("gopher") // Get a *Field based on the given field name
	fmt.Println(f1)
	//f, ok := s.FieldOk(name)  // Get a *Field based on the given field name
	//fmt.Println(f)
	//n := s.Name()             // Get the struct name
	//h := s.HasZero()          // Check if any field is uninitialized
	//z := s.IsZero()           // Check if all fields are uninitialized

}
