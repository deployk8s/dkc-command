package log

import (
	log "github.com/sirupsen/logrus"
)

type myLog struct {
	*log.Logger
}

//func (ml myLog) Println(v ...interface{}) {
//		ml.Logger.Println(v...)
//}
//func (ml myLog) Printf(format string, v ...interface{}) {
//		ml.Logger.Printf(format, v...)
//}
func NewMyLog() myLog {
	var ml myLog
	ml.Logger = log.New()
	ml.SetFormatter(&log.TextFormatter{DisableTimestamp: true})
	return ml
}

var Log myLog

func init() {
	Log = NewMyLog()
}
