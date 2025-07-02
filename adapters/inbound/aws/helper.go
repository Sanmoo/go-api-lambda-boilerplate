package aws

import (
	"log/slog"
	"os"

	"github.com/Sanmoo/go-api-lambda-boilerplate/adapters/inbound/http"
	alg "github.com/akrylysov/algnhsa"
)

func ListenAndServe(si http.ServerInterface, name string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Initializing Lambda function", "function", name)
	defer logger.Debug("Function main finished", "function", name)
	handler := http.HandlerWithOptions(si, http.StdHTTPServerOptions{
		BaseURL: "/default",
	})
	alg.ListenAndServe(handler, nil)
}
