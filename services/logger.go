package services

import (
	"errors"
	"github.com/ohoareau/gogh/common"
	"github.com/ohoareau/gogh/loggers/console"
	"github.com/ohoareau/gogh/loggers/none"
	"os"
)

var loggers = map[string]func() common.Logger{
	"none":    none.Create,
	"console": console.Create,
}

var defaultLogger = none.Create()
var logger = &defaultLogger

//goland:noinspection GoUnusedParameter
func InitLogger(options common.Options) error {
	name := "none"
	envName := os.Getenv("GH_LOGGER")
	if len(envName) > 0 {
		name = envName
	}
	if len(options.Logger) > 0 {
		name = options.Logger
	}
	f, ok := loggers[name]
	if !ok {
		return errors.New("unknown GH logger '" + name + "'")
	}
	l := f()
	logger = &l
	return nil
}

func GetLogger() *common.Logger {
	return logger
}

func RegisterLogger(name string, factory func() common.Logger) error {
	loggers[name] = factory
	return nil
}
