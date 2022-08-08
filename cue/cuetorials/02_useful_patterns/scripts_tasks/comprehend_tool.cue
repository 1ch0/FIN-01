package scripts

import (
	"tool/cli"
	"tool/file"
)

maxlen: 16

command: foreach: {
	list: file.Glob & {}
}
