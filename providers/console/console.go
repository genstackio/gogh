package console

import (
	"context"
	"fmt"
	"github.com/genstackio/gogh/common"
)

type Provider struct {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) AddCaptureContext(ctx common.CaptureContext) {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureError(err error, ctx common.CaptureContext) {
	fmt.Println(err, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessage(message string, ctx common.CaptureContext) {
	fmt.Println(message, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureMessages(messages []string, ctx common.CaptureContext) {
	fmt.Println(messages, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureProperty(typ string, data interface{}, ctx common.CaptureContext) {
	fmt.Println(typ, data, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureData(bulkData map[string]interface{}, ctx common.CaptureContext) {
	fmt.Println(bulkData, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureEvent(event interface{}, ctx common.CaptureContext) {
	fmt.Println(event, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTag(tag string, value interface{}, ctx common.CaptureContext) {
	fmt.Println(tag, value, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) CaptureTags(tags map[string]interface{}, ctx common.CaptureContext) {
	fmt.Println(tags, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Error(err error, ctx common.CaptureContext) {
	p.CaptureError(err, ctx)
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Clean() {
}

//goland:noinspection GoUnusedParameter
func (p *Provider) Wrap(h common.HandlerFn) common.HandlerFn {
	return func(ctx context.Context, payload []byte) ([]byte, error) {
		return h(ctx, payload)
	}
}

//goland:noinspection GoUnusedExportedFunction
func Create() common.Provider {
	return &Provider{}
}
