//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tomkdickinson/serverless-go-template/internal/api"
	"github.com/tomkdickinson/serverless-go-template/internal/blog"
	"github.com/tomkdickinson/serverless-go-template/internal/storage/memory"
)

func setupHandler() api.BlogPostHandlers {
	panic(wire.Build(
		memory.ProvideRepository,
		blog.ProvideService,
		api.ProvideBlogPostHandlers,
	))
}
