package common

import "context"

type CaptureContext struct {
	Tag         CaptureContextTuple
	Tags        map[string]any
	User        any
	Event       any
	Message     string
	Messages    []string
	Error       error
	Property    CaptureContextTuple
	Data        any
	Environment string
}

type CaptureContextTuple struct {
	string
	any
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
	CaptureProperty(typ string, value any, ctx CaptureContext)
	CaptureData(bulkData map[string]any, ctx CaptureContext)
	CaptureEvent(event any, ctx CaptureContext)
	CaptureTag(tag string, value any, ctx CaptureContext)
	CaptureTags(tags map[string]any, ctx CaptureContext)
	Error(err error, ctx CaptureContext)
	Wrap(HandlerFn) HandlerFn
}

type Logger interface {
	Log(level string, args ...any)
	Error(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Debug(args ...any)
}

type Enricher interface {
}
