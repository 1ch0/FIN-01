package scripts

import "tool/cli"

import "tool/file"

locallist: file.Glob & {
	glob: "*.cue"
}

command: brokenref: {
	for _, filepath in locallist.files {
		(filepath): print: cli.Print & {
			text: filepath // an inferred dependency
		}
	}
}