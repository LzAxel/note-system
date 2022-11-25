package logging

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func Init(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%v:%v", filename, f.Line),
				fmt.Sprintf("%v()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: false,
	}
	l.SetLevel(logrusLevel)

	e = logrus.NewEntry(l)
}
