//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tomkdickinson/serverless-hexagonal-go/internal/blog"
	"github.com/tomkdickinson/serverless-hexagonal-go/internal/lambda"
	"github.com/tomkdickinson/serverless-hexagonal-go/internal/storage/memory"
)

func setupHandler() lambda.BlogPostHandlers {
	panic(wire.Build(
		memory.ProvideRepository,
		blog.ProvideService,
		lambda.ProvideBlogPostHandlers,
	))
}
