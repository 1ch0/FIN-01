// 为了保证唯一性，Cue的数据不会被覆盖
// 结构

album: {
	title: string
	year:  int
	live:  bool
}

// 约束
album: {
	title: string
	year:  >1950
	live:  false
}

// 数据
album: {
	title: "Houses of the Holy"
	year:  1973
	live:  false
}
