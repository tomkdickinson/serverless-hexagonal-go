package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"os"
)

func withLogger(ctx context.Context, request events.APIGatewayProxyRequest) context.Context {
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("path", request.Path).
		Str("method", request.HTTPMethod).
		Str("request_id", request.RequestContext.RequestID).
		Logger()
	return logger.WithContext(ctx)
}
