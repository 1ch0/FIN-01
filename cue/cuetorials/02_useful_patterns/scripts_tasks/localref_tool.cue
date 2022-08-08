package scripts

import (
	"tool/cli"
	"tool/file"
)

locallist: file.Glob & {
	glob: "*.cue"
}

command: localref: {
	list: locallist

	for _, filepath in list.files {
		(filepath): print: cli.Print & {
			text: filepath
		}
	}
}