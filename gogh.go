package gogh

import (
	"context"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/genstackio/gogh/common"
	"github.com/genstackio/gogh/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GhWrapper struct {
	Handler runtime.Handler
	Options common.Options
}

func prepare(options common.Options, payload []byte) (common.CaptureContext, common.Provider) {
	c := EnrichContext(options)
	opts := common.Options{
		Logger:      options.Logger,
		Provider:    options.Provider,
		Environment: c.Environment,
	}
	localCtx := common.CaptureContext{
		Data: map[string]string{
			"payload": string(payload),
		},
	}
	err := Init(opts)
	p := *GetProvider()
	if nil != err {
		p.Error(err, localCtx)
	}

	return localCtx, p
}
func (w GhWrapper) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	localCtx, p := prepare(w.Options, payload)
	h := p.Wrap(w.Handler.Invoke)
	output, err := h(ctx, payload)
	if nil != err {
		p.Error(err, localCtx)
	}
	return output, err
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		defer (func() {
		})()
	})
}

//goland:noinspection GoUnusedExportedFunction
func GhHandlerWrapper(h runtime.Handler, options common.Options) runtime.Handler {
	return GhWrapper{
		Handler: h,
		Options: options,
	}
}

//goland:noinspection GoUnusedExportedFunction
func GhHttpRouterConfigurator(r *chi.Mux) {
	r.Use(middleware)
}

func GetProvider() *common.Provider {
	return services.GetProvider()
}
func GetLogger() *common.Logger {
	return services.GetLogger()
}

//goland:noinspection GoUnusedExportedFunction
func Init(options common.Options) error {
	err := services.InitLogger(options)
	if nil != err {
		return err
	}
	err = services.InitProvider(options)
	if nil != err {
		return err
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func EnrichContext(options common.Options) common.CaptureContext {
	return common.CaptureContext{}
}

//goland:noinspection GoUnusedExportedFunction
func Log(level string, args ...interface{}) {
	(*GetLogger()).Log(level, args...)
}

//goland:noinspection GoUnusedExportedFunction
func Error(args ...interface{}) {
	(*GetLogger()).Error(args...)
}

//goland:noinspection GoUnusedExportedFunction
func Info(args ...interface{}) {
	(*GetLogger()).Info(args...)
}

//goland:noinspection GoUnusedExportedFunction
func Warn(args ...interface{}) {
	(*GetLogger()).Warn(args...)
}

//goland:noinspection GoUnusedExportedFunction
func Debug(args ...interface{}) {
	(*GetLogger()).Debug(args...)
}

//goland:noinspection GoUnusedExportedFunction
func AddCaptureContext(ctx common.CaptureContext) {
	(*GetProvider()).AddCaptureContext(ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureError(err error, ctx common.CaptureContext) {
	(*GetProvider()).CaptureError(err, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureMessage(message string, ctx common.CaptureContext) {
	(*GetProvider()).CaptureMessage(message, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureMessages(messages []string, ctx common.CaptureContext) {
	(*GetProvider()).CaptureMessages(messages, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureProperty(typ string, data interface{}, ctx common.CaptureContext) {
	(*GetProvider()).CaptureProperty(typ, data, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureData(bulkData map[string]interface{}, ctx common.CaptureContext) {
	(*GetProvider()).CaptureData(bulkData, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureEvent(event interface{}, ctx common.CaptureContext) {
	(*GetProvider()).CaptureEvent(event, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureTag(tag string, value interface{}, ctx common.CaptureContext) {
	(*GetProvider()).CaptureTag(tag, value, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func CaptureTags(tags map[string]interface{}, ctx common.CaptureContext) {
	(*GetProvider()).CaptureTags(tags, ctx)
}

//goland:noinspection GoUnusedExportedFunction
func RegisterProvider(name string, factory func() common.Provider) error {
	return services.RegisterProvider(name, factory)
}

//goland:noinspection GoUnusedExportedFunction
func RegisterLogger(name string, factory func() common.Logger) error {
	return services.RegisterLogger(name, factory)
}

//goland:noinspection GoUnusedExportedFunction
func RegisterEnricher(name string, factory func() common.Enricher) error {
	return services.RegisterEnricher(name, factory)
}
