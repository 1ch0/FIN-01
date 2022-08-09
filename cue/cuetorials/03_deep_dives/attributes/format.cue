foo: "bar" @attr()

foo: {
	@attr()
	bar: string
}

any: _ @attr(foo, bar)

any: _ @attr(key1, key2=value, key3="foo;bar")