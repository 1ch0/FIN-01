package scripts

import "tool/exec"

import "tool/file"

command: enforced: {
	mkdir: file.MkdirAll & {
		path: "./out/forced/write"
	}

	write: file.Create & {
		$dep:     mkdir.$done //explicit dependency
		filename: "./out/forced/write/foo.txt"
		contents: "========="
	}

	clean: exec.Run & {
		$dep: write.$done
		cmd:  "rm -rf ./out"
	}
}
