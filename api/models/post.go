package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID  `jsonapi:"primary,posts`
	Title     string     `jsonapi:"attr,title"`
	AuthorID  uuid.UUID  `jsonapi:"attr,author_id"`
	CreatedAt time.Time  `jsonapi:"attr,created_at"`
	ViewCount int        `jsonapi:"attr,view_count"`
	Comments  []*Comment `jsonapi:"relation,comments"`
}
