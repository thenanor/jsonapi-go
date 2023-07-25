package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID  `json:"id" jsonapi:"primary,posts"`
	Title     string     `json:"title" jsonapi:"attr,title"`
	AuthorID  uuid.UUID  `json:"author_id" jsonapi:"attr,author_id"`
	CreatedAt time.Time  `json:"created_at" jsonapi:"attr,created_at"`
	ViewCount int        `json:"view_count" jsonapi:"attr,view_count"`
	Comments  []*Comment `json:"comments" jsonapi:"relation,comments"`
}
