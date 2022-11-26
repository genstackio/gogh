package gogh

import (
	"context"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type GhWrapper struct {
	Handler runtime.Handler
}

func (w GhWrapper) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	log.Println(string(payload))
	return w.Handler.Invoke(ctx, payload)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Body)
		next.ServeHTTP(w, r)
	})
}

func GhHandlerWrapper(h runtime.Handler) runtime.Handler {
	return GhWrapper{Handler: h}
}

func GhHttpRouterConfigurator(r *chi.Mux) {
	r.Use(middleware)
}
