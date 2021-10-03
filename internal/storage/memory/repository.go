package memory

import (
	"context"
	"errors"
	"github.com/tomkdickinson/serverless-hexagonal-go/internal/blog"
	"sort"
	"time"
)

type BlogDatabase map[string]Entry

func ProvideRepository() blog.Repository {
	return BlogDatabase{
		"hello": Entry{
			Title:   "Hello",
			Content: "Hello...",
			Author:  "Tom Dickinson",
			Time:    time.Date(2021, time.October, 2, 15, 50, 0, 0, time.UTC),
		},
		"world": Entry{
			Title:   "World!",
			Content: "... World!",
			Author:  "Tom Dickinson",
			Time:    time.Date(2021, time.October, 2, 15, 55, 0, 0, time.UTC),
		},
	}
}

func (b BlogDatabase) FindOne(ctx context.Context, slug string) (*blog.Entry, error) {
	if slug == "" {
		return nil, errors.New("slug cannot be empty")
	}
	if entry, exists := b[slug]; exists {
		return entry.toDomainModel(slug), nil
	}
	return nil, nil
}

func (b BlogDatabase) Find(ctx context.Context) (*[]blog.Entry, error) {
	var entries []blog.Entry

	for slug, entry := range b {
		entries = append(entries, *entry.toDomainModel(slug))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].CreatedAt.UnixNano() < entries[j].CreatedAt.UnixNano()
	})

	return &entries, nil
}
