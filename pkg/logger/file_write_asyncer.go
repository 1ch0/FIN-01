package logger

import (
	"bytes"
	"errors"
	"time"

	"github.com/natefinch/lumberjack"
)

type FileWriteAsyncer struct {
	innerLogger *lumberjack.Logger
	ch          chan []byte
	syncChan    chan struct{}
}

func NewFileWriteAsyncer(filepath string) *FileWriteAsyncer {
	fa := &FileWriteAsyncer{}
	fa.innerLogger = &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	fa.ch = make(chan []byte, 10000)
	fa.syncChan = make(chan struct{})
	go batchWriteLog(fa)
	return fa

}

func (fa *FileWriteAsyncer) Write(data []byte) (int, error) {
	select {
	case fa.ch <- data:
		return len(data), nil
	default:
		// 注意: 需要在metric中记录丢弃的日志数量
		return 0, errors.New("Log channel is full")
	}
}

func (fa *FileWriteAsyncer) Sync() error {
	fa.syncChan <- struct{}{}
	return nil
}

func batchWriteLog(fa *FileWriteAsyncer) {
	buffer := bytes.NewBuffer(make([]byte, 0, 10240))

	ticker := time.NewTicker(time.Millisecond * 200)
	//var record []byte
	var err error
	for {
		select {
		case <-ticker.C:
			if len(buffer.Bytes()) > 0 {
				_, err = fa.innerLogger.Write(buffer.Bytes())
				if err != nil {
					panic(err)
				}
				buffer.Reset()
			}

		case record := <-fa.ch:
			buffer.Write(record)
			if len(buffer.Bytes()) >= 1024*4 {
				_, err = fa.innerLogger.Write(buffer.Bytes())
				if err != nil {
					panic(err)
				}
				buffer.Reset()
			}
		case <-fa.syncChan:
			if len(buffer.Bytes()) > 0 {
				_, err = fa.innerLogger.Write(buffer.Bytes())
				if err != nil {
					panic(err)
				}
				buffer.Reset()
			}
			break
		}
	}

}
