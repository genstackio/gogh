package services

import (
	"errors"
	"github.com/genstackio/gogh/common"
	"github.com/genstackio/gogh/providers/console"
	"github.com/genstackio/gogh/providers/none"
	"os"
)

var providers = map[string]func() common.Provider{
	"none":    none.Create,
	"console": console.Create,
}

var defaultProvider = none.Create()
var provider = &defaultProvider

//goland:noinspection GoUnusedParameter
func InitProvider(options common.Options) error {
	name := "console"
	envName := os.Getenv("GH_PROVIDER")
	if len(envName) > 0 {
		name = envName
	}
	if len(options.Provider) > 0 {
		name = options.Provider
	}
	f, ok := providers[name]
	if !ok {
		return errors.New("unknown GH provider '" + name + "'")
	}
	l := f()
	provider = &l
	return nil
}

func GetProvider() *common.Provider {
	return provider
}

func RegisterProvider(name string, factory func() common.Provider) error {
	providers[name] = factory
	return nil
}
