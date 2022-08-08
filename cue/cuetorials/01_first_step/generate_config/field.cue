import "strings"


// list 推导： [ for key, val in iterable { ... } ]，可以不使用 key，只用 value。
//field 推导：[ for key, val in iterable { ... }]，我们调用内置的函数插入一个 field，注意要用双引号包裹。
//字符串插入："(<a cue expression)"`，这是基于 Swift 的字符串插入，而且当你想要兼容 JSON 时，这个是唯一的机制。
//隐藏字段：_hidden: "I'm hidden" 以下划线开头，CUE 也有隐藏定义 （_#）
band: {
	name: "Led Zeppelin"

	selfTitled: [for i, I in _selfIndexAlbums {title: "\(name) \(I)"}]

	allAlbums: [
		for I in _selfIndexAlbums {title: "\(name) \(I)"},
		for N in _nameAlbums {title: "\(N)"},
	]

	Albums: [
		for key, val in allAlbums {
			"\(strings.TrimSpace(val.title))": {
				pos: key,
				aritst: name,
				title: strings.TrimSpace(val.title),
				titleLen: len(val.title),
			}
		}
	]

	_selfIndexAlbums: ["", "II", "III", "IV"]
	_nameAlbums: [
		"Houses of the Holy",
		"Physical Graffiti",
		"Presence",
		"In Through the Out Door",
		"Coda",
	]
}