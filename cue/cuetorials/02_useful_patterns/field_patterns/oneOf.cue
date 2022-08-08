#OneOf: {a: string} | {b: string}

// #OneOf: {} | {a: string} | {b: string} | {c: int} | {d: bool}

// you cannot have options which are subsets of another option.
// The following will error when only a is provided
// #OneOf: {a: _} | {a: _, b: _}

#E: {
	name: string
	#OneOf
}

ex1: #E & {
	name: "a choice"
	a:    "bar"
}

ex2: #E & {
	name: "b choice"
	b:    "hello"
}

ex3: #E & {
	name: "error none chosen"
}

ex4: #E & {
	name: "error both chosen"
	a:    "a"
	b:    "b"
}
