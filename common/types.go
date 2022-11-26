package common

import "context"

type CaptureContext struct {
	Tag         CaptureContextTuple
	Tags        map[string]interface{}
	User        interface{}
	Event       interface{}
	Message     string
	Messages    []string
	Error       error
	Property    CaptureContextTuple
	Data        interface{}
	Environment string
}

type CaptureContextTuple struct {
	Key   string
	Value interface{}
}

type Options struct {
	Logger      string
	Provider    string
	Environment string
}

type HandlerFn func(ctx context.Context, payload []byte) ([]byte, error)

type Provider interface {
	AddCaptureContext(ctx CaptureContext)
	CaptureError(err error, ctx CaptureContext)
	CaptureMessage(message string, ctx CaptureContext)
	CaptureMessages(messages []string, ctx CaptureContext)
	CaptureProperty(typ string, value interface{}, ctx CaptureContext)
	CaptureData(bulkData map[string]interface{}, ctx CaptureContext)
	CaptureEvent(event interface{}, ctx CaptureContext)
	CaptureTag(tag string, value interface{}, ctx CaptureContext)
	CaptureTags(tags map[string]interface{}, ctx CaptureContext)
	Error(err error, ctx CaptureContext)
	Wrap(HandlerFn) HandlerFn
}

type Logger interface {
	Log(level string, args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
}

type Enricher interface {
}
