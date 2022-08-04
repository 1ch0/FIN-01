// cuelang 用*表示默认值：
// 缺省值
test: {
	name: *parameter.name | "test"
	value: *parameter.value | 0
}

parameter: {
	name: string
	value?: int
}

parameter: {
	name: "new"
}