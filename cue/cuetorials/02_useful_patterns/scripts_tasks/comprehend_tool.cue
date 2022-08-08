package scripts

import (
	"tool/cli"
	"tool/file"
)

maxlen: 16

command: foreach: {
	list: file.Glob & {
		glob: "*.cue"
	}

	for _, filepath in list.files {
		(filepath): {
			read: file.Read & {
				filename: filepath
				contents: string
			}
			print: cli.Print & {
				text: read.contents
			}
		}
	}
}

command: maybe: {
	list: file.Glob & {
		glob: "*.cue"
	}

	for _, filepath in list.files {
		if len(filepath) > maxlen {
			(filepath): print: cli.Print & {
				text: filepath
			}
		}
	}
}