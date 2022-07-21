package log

import (
	"go.uber.org/zap"
)

var ZapLogger *zap.SugaredLogger

func init() {
	l, _ := zap.NewProduction()
	ZapLogger = l.Sugar()
}
