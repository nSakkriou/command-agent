package util

import (
	"fmt"
	"time"

	"github.com/theritikchoure/logx"
)

func Info(message string, format ...interface{}) {
	logfLevel("INFO", message, logx.FGGREEN, format...)
}

func Debug(message string, format ...interface{}) {
	logfLevel("DEBUG", message, logx.FGBLUE, format...)
}

func Raw(message string, format ...interface{}) {
	logf(message, "", format...)
}

func Error(message string, format ...interface{}) {
	logfLevel("ERROR", message, logx.FGRED, format...)
}

func Verbose(message string, format ...interface{}) {
	logfLevel("VERBOSE", message, logx.FGYELLOW, format...)
}

// Ne pas utiliser
func logf(message, color string, format ...interface{}) {
	if len(format) == 0 {
		logx.Logf(message, color, "", format...)
	} else {
		formattedMessage := fmt.Sprintf(message, format...)
		logx.Logf(formattedMessage, color, "")
	}
}

func logfLevel(level, message, color string, format ...interface{}) {
	prefixedMessage := fmt.Sprintf("[%s] [%s] %s", timestamp(), level, message)
	logf(prefixedMessage, color, format...)
}

func timestamp() string {
	now := time.Now()
	return now.Format("15:04:05")
}
