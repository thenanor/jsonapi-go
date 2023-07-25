package handlers

import (
	"context"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/thenanor/jsonapi-go/api/models"
)

func (router *Router) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}

	// Read the post from body and convert it to struct
	if err := jsonapi.UnmarshalPayload(r.Body, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate in BL and save in DL
	postInternal, err := router.bl.CreatePost(context.Background(), post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	if err := jsonapi.MarshalPayload(w, postInternal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
