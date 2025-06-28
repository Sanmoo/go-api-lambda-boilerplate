package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/Sanmoo/go-api-lambda-boilerplate/core/usecases"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaHandler struct {
	logger *slog.Logger
}

func (h *LambdaHandler) handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	h.logger.Info("Request", "request", request)
	h.logger.Info("Context", "context", ctx)
	defer h.logger.Info("Request handling finished", "request")
	
	if request.HTTPMethod == "GET" {
		books, _ := usecases.ListBooks()
		responseJson, err := json.Marshal(books)
	} else if request.HTTPMethod == "POST" {
		h.logger.Error("POST method is not implemented yet")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       ,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Initializing Lambda function", "function", "books")
	defer logger.Info("Function main finished", "function", "books")

	lambdaHandler := &LambdaHandler{
		logger: logger,
	}

	lambda.Start(lambdaHandler.handleRequest)
}
