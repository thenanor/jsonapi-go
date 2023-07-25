package models

type Comment struct {
	ID     string `json:"id" jsonapi:"primary,comments"`
	PostID string `json:"post_id" jsonapi:"attr,post_id"`
	Body   string `json:"body" jsonapi:"attr,body"`
	Likes  uint   `json:"likes_count" jsonapi:"attr,likes_count,omitempty"`
}
