import "list"

val: this={
	a?:        string
	b?:        string
	c?:        string
	#AnyOfABC: true & list.MinItems([ for label, _ in this if list.Contains(["a", "b", "c"], label) {label}], 2)

	d?:        string
	e?:        string
	#AnyOfABC: true & list.MinItems([ for label, _ in this if list.Contains(["d", "e"], label) {label}], 1)

	#AnyOfXYZ & {#CheckXYZ: true}

	a: "a"
	b: "bbb"
	d: "d"
	z: "z"
}

#AnyOfXYZ: this={
	x?:        string
	y?:        string
	z?:        string
	#CheckXYZ: list.MinItems([ for label, _ in this if list.Contains(["x", "y", "z"], label) {label}], 1)
}
