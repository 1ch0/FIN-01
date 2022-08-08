package function

import "strings"

#Input: {
	count: int
	msg:   string
}

#Output: {
	val: string
}

#Transform: {
	X1="in": #Input
	out:     #Output

	_upper: strings.ToUpper(X1.msg)
	_msg:   strings.Join([_upper]*X1.count, " ")

	out: val: _msg
}

result: #Transform & {in: {msg: "ra", count: 3}}
