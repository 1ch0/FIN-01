// 默认情况下，struct 是开放的，而 definition 是封闭的。
// Closed struct
s: close({
	foo: "bar"
})

// Open definition
#d: {
	foo: "bar"
	... // must be last
}
