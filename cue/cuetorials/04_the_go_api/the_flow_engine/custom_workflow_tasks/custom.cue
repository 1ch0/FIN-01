Custom Flow111 Task
TF:  {
	a: {
		foo:   1
		hello: string
	}
	b: {
		foo: 2
	}
	c: {
		foo: 3
		goo: 6
	}
}
TF:  {
	foo:   1
	hello: string
}
TF:  {
	foo: 2
}
TF:  {
	foo: 3
	goo: 6
}
CustomTask: 2 {
	foo: 2
}
CustomTask: 1 {
	foo:   1
	hello: string
}
TF:  {
	a: {
		foo:   1
		hello: string
	}
	b: {
		foo:   2
		bar:   3
		hello: "world"
	}
	c: {
		foo: 3
		goo: 6
	}
}
TF:  {
	a: {
		foo:   1
		bar:   2
		hello: "world"
	}
	b: {
		foo:   2
		bar:   3
		hello: "world"
	}
	c: {
		foo: 3
		goo: 6
	}
}
CustomTask: 3 {
	foo: 3
	goo: 6
}
TF:  {
	a: {
		foo:   1
		bar:   2
		hello: "world"
	}
	b: {
		foo:   2
		bar:   3
		hello: "world"
	}
	c: {
		foo:   3
		bar:   4
		goo:   6
		hello: "world"
	}
}
