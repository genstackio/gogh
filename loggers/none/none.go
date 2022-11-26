package none

import "github.com/ohoareau/gogh/common"

type Logger struct {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Log(level string, args ...interface{}) {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Error(args ...interface{}) {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Info(args ...interface{}) {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Warn(args ...interface{}) {
}

//goland:noinspection GoUnusedParameter
func (l *Logger) Debug(args ...interface{}) {
}

//goland:noinspection GoUnusedExportedFunction
func Create() common.Logger {
	return &Logger{}
}
