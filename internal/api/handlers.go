package api

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomkdickinson/serverless-go-template/internal/blog"
	"net/http"
	"os"
)

type Response events.APIGatewayProxyResponse

type BlogPostHandlers struct {
	svc blog.Service
}

func ProvideBlogPostHandlers(svc blog.Service) BlogPostHandlers {
	return BlogPostHandlers{
		svc: svc,
	}
}

func (h BlogPostHandlers) GetPost(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	ctx = withLogger(ctx, request)
	entries, err := h.svc.GetEntry(ctx, request.PathParameters["slug"])

	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not load entry")
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	if entries == nil {
		log.Ctx(ctx).Info().Err(err).Msg("Entry not found")
		return Response{StatusCode: http.StatusNotFound}, nil
	}

	body, err := json.Marshal(newEntryPayload(*entries))
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not marshal entry to JSON")
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func (h BlogPostHandlers) ListPosts(ctx context.Context) (Response, error) {
	entries, err := h.svc.ListEntries(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not load entries")
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	body, err := json.Marshal(multipleEntries(*entries))
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not marshal entry to JSON")
		return Response{StatusCode: http.StatusInternalServerError}, err
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

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