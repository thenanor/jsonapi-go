package models

import (
	"github.com/google/uuid"
)

type Author struct {
	ID    uuid.UUID `jsonapi:"primary,authors"`
	Name  string    `jsonapi:"attr,name"`
	Posts []*Post   `jsonapi:relation,posts`
}
