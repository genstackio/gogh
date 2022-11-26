package services

import "github.com/ohoareau/gogh/common"

var enrichers = map[string]func() common.Enricher{}

func RegisterEnricher(name string, factory func() common.Enricher) error {
	enrichers[name] = factory
	return nil
}
