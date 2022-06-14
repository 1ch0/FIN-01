package log

import (
	"fmt"
	"io"
	"os"
)

type Level uint8

const (
	Debuglevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[Level]string{
	Debuglevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

type options struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     fmt.Formatter
	disableCaller bool
}

type Option func(*options)

func initOptinos(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}
