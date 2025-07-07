package aws

import (
	"log/slog"
	"os"

	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	alg "github.com/akrylysov/algnhsa"
	"github.com/getkin/kin-openapi/openapi3"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func ListenAndServe(si http.ServerInterface, name string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Initializing Lambda function", "function", name)
	defer logger.Debug("Function main finished", "function", name)

	spec, err := openapi3.NewLoader().LoadFromData([]byte(http.OpenAPI))
	if err != nil {
		panic(err)
	}
	spec.Servers = nil

	mw := middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{})

	handler := http.HandlerWithOptions(si, http.StdHTTPServerOptions{
		Middlewares: []http.MiddlewareFunc{mw},
	})

	alg.ListenAndServe(handler, &alg.Options{
		DebugLog: true,
	})
}
