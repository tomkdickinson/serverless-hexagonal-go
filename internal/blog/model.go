package blog

import "time"

type Entry struct {
	Title     string
	Slug      string
	Content   string
	Author    string
	CreatedAt time.Time
}
