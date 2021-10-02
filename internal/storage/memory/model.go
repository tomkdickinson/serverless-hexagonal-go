package memory

import (
	"github.com/tomkdickinson/serverless-go-template/internal/blog"
	"time"
)

type Entry struct {
	Title   string
	Content string
	Author  string
	Time    time.Time
}

func (e Entry) toDomainModel(slug string) *blog.Entry {
	return &blog.Entry{
		Title:     e.Title,
		Slug:      slug,
		Content:   e.Content,
		Author:    e.Author,
		CreatedAt: e.Time,
	}
}
