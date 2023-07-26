package models

type Author struct {
	ID   string `json:"id" jsonapi:"primary,authors"`
	Name string `json:"name" jsonapi:"attr,name"`
	// Posts []*Post `json:"posts" jsonapi:"relation,posts"`
}
