container: {
	repo: "docker.io/cuelang"
	image: "cue"
	version: "v0.3.0"
	full: "\(repo)/\(image):\(version)"
}

name: "Tony"
msg: "Hello, \(name)"
// convert string to bytes
b: '\(msg)'
// convert bytes to string
s:"\(b)"