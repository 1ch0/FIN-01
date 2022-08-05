// 隐式字段和值通过在前面增加下划线定义，可以在当前 package 进行引用，
// 在计算时可以可选的显示它们，cue eval -H hidden.cue
// 隐式字段在不同 package 中是不可见的，所以不可以使用引用的 package 中隐式字段
A: {
	_hidden: "a hidden field"
	isshown: "I can be seen"
	hidrefd: _hidden + " sort of?"
}

_#NoshowDefn: {
	hello: string
	num:   int | *42
}

B: _#NoshowDefn & {hello: "world"}
