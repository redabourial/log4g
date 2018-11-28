package log4g

import (
	"fmt"
	"os"
	"time"
)

func NewConsoleLogger() LoggerStream {
	return func(level string, values ...interface{}) {
		var file *os.File
		if level == FATAL || level == ERROR {
			file = os.Stderr
		} else {
			file = os.Stdout
		}
		fmt.Fprintf(file, "%s %s : %v\r\n", time.Now().Format(time.RFC1123), level, values)
	}
}