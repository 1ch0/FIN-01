// disjunction os values (like an enum)
// 使用|可以实现支持多种结构。同时它也可以为出错值设置替换值
hello: "world" | "bob" | "mary"
hello: "world"

// disjunction of types
port: string | int
port: 5432

// disjunction of schemas
val: #Def1 | #Def2
val: {foo: "bar", ans: 42}

#Def1: {
	foo: string
	ans: int
}

#Def2: {
	name: string
	port: int
}
