package console

import (
	"fmt"
	"github.com/ohoareau/gogh/common"
)

type Logger struct {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Log(level string, args ...interface{}) {
	a := []interface{}{
		"[" + level + "]",
	}
	for i := 0; i < len(args); i++ {
		a = append(a, args[i])
	}
	fmt.Println(args...)
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Error(args ...interface{}) {
	l.Log("error", args...)
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Info(args ...interface{}) {
	l.Log("info", args...)
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Warn(args ...interface{}) {
	l.Log("warn", args...)
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Debug(args ...interface{}) {
	l.Log("debug", args...)
}

//goland:noinspection GoUnusedExportedFunction
func Create() common.Logger {
	return &Logger{}
}
