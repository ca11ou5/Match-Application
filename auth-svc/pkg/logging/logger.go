package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type Logger struct {
	*logrus.Entry
}

var entry *logrus.Entry

func InitLogger() *Logger {
	return &Logger{entry}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "03:04:05 02/01/2006",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return fmt.Sprintf("%s()", path.Base(frame.Function)), fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
		},
	}
	l.SetOutput(io.Discard)
	l.SetLevel(logrus.DebugLevel)
	l.AddHook(&writerHook{
		writer:    os.Stdout,
		logLevels: logrus.AllLevels,
	})

	entry = logrus.NewEntry(l)
}

type writerHook struct {
	writer    io.Writer
	logLevels []logrus.Level
}

func (wh *writerHook) Levels() []logrus.Level {
	return wh.logLevels
}

func (wh *writerHook) Fire(e *logrus.Entry) error {
	line, err := e.String()
	if err != nil {
		return err
	}

	wh.writer.Write([]byte(line))
	return err
}
