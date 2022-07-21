package logger

import (
	"errors"
	"testing"
)

func TestJSONLogger(t *testing.T) {
	logger, err := NewJSONLogger(
		WithEnableConsole(true),
		//WithField("domain", fmt.Sprintf("%s", config.ProjectName)),
		//WithFileP(config.ProjectLogFile),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer logger.Sync()

	err = errors.New("pkg error")
	logger.Error("err occurs", WrapMeta(nil, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)
	logger.Error("err occurs", WrapMeta(err, NewMeta("para1", "value1"), NewMeta("para2", "value2"))...)

}
