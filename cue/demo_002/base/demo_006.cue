// 通过定义的最后添加...来申明开放模式定义；
// 另外通过过close强制为结构体设置为关闭模式
#d: {
	foo: "bar"
	... // must be last
}

// Closed struct
s: close({
	foo: "bar"
})

jim: {
	name: "Jim"
}

jim: {
	age: 12
}