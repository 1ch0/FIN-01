test: {
	name: parameter["name"]
	value: parameter["value"]
}

parameter: {
	name: [...string]
	value?: [... >=0 & < 10 & int]
}

parameter: {
	name: ["test1", "test2"]
	value: [5, 9]
}

test2: {
	key: [
		for k, v in parameter2["name"]{
			name: v
		}
	]
}

parameter2: {
	name: [...string]
}

parameter2: {
	name: [
		"job1",
		"job2"
		]
}