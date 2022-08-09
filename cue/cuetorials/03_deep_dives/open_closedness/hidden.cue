
#D: {
	size: string
	data: {
		x: int
		y: int
	}
}

// after conjuncting, extension is not allowed
// You can add hidden fields to a closed value.
// This works for both definitions and structs which have been close()’d.
d: #D & {
	data: {
		x: 3
		y: 4
	}

	_calc: data.x * data.y
	size:  string | *"med"
	if _calc < 10 {
		size: "small"
	}
	if _calc > 100 {
		size: "large"
	}
}