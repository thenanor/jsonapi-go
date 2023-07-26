package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/jsonapi"
	"github.com/thenanor/jsonapi-go/businesslogic"
)

const (
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
)

type Router struct {
	bl businesslogic.Businesslogic
}

func New() (*Router, error) {
	bl, err := businesslogic.New()
	if err != nil {
		return nil, fmt.Errorf("unable to create businesslogic: %w", err)
	}

	return &Router{
		bl,
	}, nil
}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(HeaderAccept) != jsonapi.MediaType {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
	}

	var methodHandler http.HandlerFunc
	switch r.Method {
	case http.MethodPost:
		methodHandler = h.CreatePost
	// case http.MethodPut:
	// 	methodHandler = h.UpdatePosts
	case http.MethodGet:
		if strings.TrimPrefix(r.URL.Path, "/posts/") != "" {
			methodHandler = h.GetPost
		} else {
			methodHandler = h.GetPosts
		}
	default:
		w.Header().Set("Content-Type", jsonapi.MediaType)
		http.Error(w, "Method Not Found", http.StatusNotFound)
		return
	}

	methodHandler(w, r)
}
