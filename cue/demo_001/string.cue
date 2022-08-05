import (
	"strconv"
	"encoding/json"
)

test: {
	name: strconv.FormatInt(parameter.value, 10)
	test: "example: \(json.Marshal(testvalue))"
}

let testvalue = {"value": 10}

parameter: {
	name:  string
	value: int
}

parameter: {
	name:  "test"
	value: 10
}
