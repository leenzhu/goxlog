package xlog

import(
	"fmt"
	"log"
	"os"
	"github.com/leenzhu/goutils"
)

type XLogger struct {
	stdOn bool
	name string

	file *log.Logger
	stdOut *log.Logger
}
var l XLogger

type logLevel int
const (
	DEBUG logLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)
func (l logLevel) String() string {
	tagName := map[logLevel]string{
		DEBUG : "DEBUG",
		INFO : "INFO",
		WARN : "WARN",
		ERROR : "ERROR",
		FATAL : "FATAL",
	}
	return tagName[l]
}

func init() {
	l.stdOn = true
	l.name = fmt.Sprintf("%s.log", utils.GetProcessName())

	f, err := os.OpenFile(l.name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		return
	}

	l.file = log.New(f, "", log.LstdFlags|log.Lshortfile)
	l.stdOut = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
}

func Debugf(format string, v ...interface{}) {
	l.output(DEBUG, format, v...)
}

func Infof(format string, v ...interface{}) {
	l.output(INFO, format, v...)
}

func Warnf(format string, v ...interface{}) {
	l.output(WARN, format, v...)
}

func Errorf(format string, v ...interface{}) {
	l.output(ERROR, format, v...)
}

func Fatalf(format string, v ...interface{}) {
	l.output(FATAL, format, v...)
}

func (this *XLogger) output(level logLevel, format string, v ...interface{}) {
	format = fmt.Sprintf("%5s %s", level, format)
	output := fmt.Sprintf(format, v...)
	if this.stdOn {
		l.stdOut.Output(3, output)
	}

	if this.file != nil {
		this.file.Output(3, output)
	}
}

