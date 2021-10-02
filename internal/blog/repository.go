package blog

import (
	"context"
	"errors"
)

var ErrRepository = errors.New("post repository execution error")

type Repository interface {
	FindOne(ctx context.Context, slug string) (*Entry, error)
	Find(ctx context.Context) (*[]Entry, error)
}
