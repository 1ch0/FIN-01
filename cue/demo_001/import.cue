import (
	"strconv"
	"encoding/json"
)

test: {
	name: strconv.FormatInt(parameter.value, 10)
	test: json.Marshal(testvalue)
}

let testvalue = {"value": 10}

parameter: {
	name:  string
	value: int
}

// 设置值
parameter: {
	name:  "test"
	value: 10
}
