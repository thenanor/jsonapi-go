package datalayer

import (
	"context"
	"fmt"
	"os"

	"github.com/thenanor/jsonapi-go/api/models"
	"github.com/thenanor/jsonapi-go/datalayer/memory"
)

type Datalayer interface {
	CreatePost(context.Context, models.Post) error

	CreateAuthor(context.Context, models.Author) error
}

func New() (Datalayer, error) {
	dlType, exists := os.LookupEnv("DL_TYPE")
	if !exists {
		dlType = "memory"
	}

	switch dlType {
	case "memory":
		return memory.New()
		// case "redis":
		// 	return redis.New()
	}

	return nil, fmt.Errorf("unknown datalayer type: %s", dlType)
}
