package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/theritikchoure/logx"
)

func Error(message string, format ...interface{}) {
	logfLevel("ERROR", message, logx.FGRED, format...)
}

func Info(message string, format ...interface{}) {
	logfLevel("INFO", message, logx.FGGREEN, format...)
}

func Verbose(message string, format ...interface{}) {
	logfLevel("VERBOSE", message, logx.FGYELLOW, format...)
}

func Debug(message string, format ...interface{}) {
	logfLevel("DEBUG", message, logx.FGBLUE, format...)
}

func Raw(message string, format ...interface{}) {
	logf(message, "", format...)
}

// Ne pas utiliser directement
func logf(message, color string, format ...interface{}) {
	message = strings.ToLower(message)

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
