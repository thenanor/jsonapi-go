package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/jsonapi"
	"github.com/thenanor/jsonapi-go/api/models"
)

func (router *Router) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := new(models.Post)

	// Read the post from body and convert it to struct
	if err := jsonapi.UnmarshalPayload(r.Body, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postInternal, err := router.bl.CreatePost(context.Background(), *post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	if err := jsonapi.MarshalPayload(w, &postInternal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (router *Router) GetPost(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")

	post, err := router.bl.GetPost(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	queryParams := r.URL.Query()
	fmt.Println("queryParams=", queryParams)

	include := queryParams.Get("include")
	if include == "" {
		if err := jsonapi.MarshalPayloadWithoutIncluded(w, &post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		includes := strings.Split(include, ",")
		fmt.Println("length of includes", len(includes))
		payload, err := jsonapi.Marshal(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// clear the not included includes from the payload

		//write to the response
		json.NewEncoder(w).Encode(payload)
	}
}

func (router *Router) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := router.bl.GetPosts(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, &posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
