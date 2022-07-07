package main

import (
	"encoding/json"
	"fmt"
)

type jsoninput []struct {
	Data string `json:"data"`
}

func main() {
	resp := `[{"data":"some data"}, {"data":"some more data"}]`
	data := &jsoninput{}
	_ = json.Unmarshal([]byte(resp), data)
	for _, value := range *data {
		fmt.Println(value.Data) // Prints "some data" and "some more data"
	}
}
