// Cue核心规则
// 1. 数据可被重复定义，但必须值保持一致
// 2. 结构字段可以被更强限制覆盖
// 3. 结构的字段会被合并，如果是列表，必须严格匹配
// 4. 规则可被递规应用
hello: "world"
hello: "world"

// set a type
s: {a: int}

// set some data
s: { a: 1, b: 2}

// set a nested field without curly braces
s: c: d: 3

// lists must have the same element
// and cannot change length
l: ["abc", "123"]
l: [
	"abc",
	"123"
]