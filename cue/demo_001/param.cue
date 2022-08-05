import (
	"encoding/json"
)

test: {
	name:  parameter.name
	value: json.Marshal(testvalue)
}

let testvalue = {"value": 10}

// 参数为可选项 ?
parameter: {
	name:   string
	value?: int
}

parameter: {
	name: "test"
}
