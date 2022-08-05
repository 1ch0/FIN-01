test: {
	name: parameter["name"]
	//  _|_ 代表空对象
	if parameter.value != _|_ {
		let value = parameter.value
	}
}

parameter: {
	name:   string
	value?: int
}

parameter: {
	name: "test"
}
