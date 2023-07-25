package models

import (
	"github.com/google/uuid"
)

type Author struct {
	ID    uuid.UUID `json:"id" jsonapi:"primary,authors"`
	Name  string    `json:"name" jsonapi:"attr,name"`
	Posts []*Post   `json:"posts" jsonapi:"relation,posts"`
}
