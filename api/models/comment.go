package models

import "github.com/google/uuid"

type Comment struct {
	ID     uuid.UUID `jsonapi:"primary,comments"`
	PostID uuid.UUID `jsonapi:"attr,post_id"`
	Body   string    `jsonapi:"attr,body"`
	Likes  uint      `jsonapi:"attr,likes_count,omitempty"`
}
