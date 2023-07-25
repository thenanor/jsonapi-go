package models

import "time"

type Post struct {
	ID    string `json:"id" jsonapi:"primary,posts"`
	Title string `json:"title" jsonapi:"attr,title"`
	// AuthorID  string    `json:"author_id" jsonapi:"attr,author_id"`
	CreatedAt time.Time `json:"created_at" jsonapi:"attr,created_at"`
	ViewCount int       `json:"view_count" jsonapi:"attr,view_count"`
	// Comments  []*Comment `json:"comments" jsonapi:"relation,comments"`
}
