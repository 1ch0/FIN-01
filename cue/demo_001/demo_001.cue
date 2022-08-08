test: {
	name:  parameter["name"]
	value: parameter["value"]
}

parameter: {
	name: string
	//       逻辑运算符 |  &
	value?: *0 | >=0 & <10
}

parameter: {
	name: "test"
	//     value: 10
	value: 5
}
