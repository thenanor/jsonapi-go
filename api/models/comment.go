package models

import "github.com/google/uuid"

type Comment struct {
	ID     uuid.UUID `json:"id" jsonapi:"primary,comments"`
	PostID uuid.UUID `json:"post_id" jsonapi:"attr,post_id"`
	Body   string    `json:"body" jsonapi:"attr,body"`
	Likes  uint      `json:"likes_count" jsonapi:"attr,likes_count,omitempty"`
}
