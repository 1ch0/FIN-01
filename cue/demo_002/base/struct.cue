#Album: {
	artist: string
	title: string
	year: int

	// ... uncomment to open, must be last
}

// This is a conjunction, it says "album" has to be "#Album"
album: #Album & {
	artist: "Led Zeppelin"
	title: "Led Zeppelin I"
	year: 1999

	// studio: true (uncomment to trigger error)
}

#Person: {
	name: string
	... // open struct
}

Jim: #Person & {
	name: "Jill"
	age: 18
}

a: {
	foo: "bar"
}

a: hello: "world"

b: close({
	left: "right"
})