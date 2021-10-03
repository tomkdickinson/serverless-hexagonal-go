package lambda

import "github.com/tomkdickinson/serverless-hexagonal-go/internal/blog"

type EntryPayload struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func newEntryPayload(domainEntry blog.Entry) EntryPayload {
	return EntryPayload{
		Title:   domainEntry.Title,
		Slug:    domainEntry.Slug,
		Content: domainEntry.Content,
		Author:  domainEntry.Author,
	}
}

func multipleEntries(domainEntries []blog.Entry) []EntryPayload {
	var entries []EntryPayload
	for _, domainEntry := range domainEntries {
		entries = append(entries, newEntryPayload(domainEntry))
	}
	return entries
}
