package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/require"
	"github.com/thenanor/jsonapi-go/api/models"
	"github.com/thenanor/jsonapi-go/utils"
)

func TestCreatePost(t *testing.T) {
	post := utils.FixturePost()

	in := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in, post); err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/posts", in)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	router, _ := New()
	result := httptest.NewRecorder()

	//function under test
	router.CreatePost(result, req)

	require.Equal(t, http.StatusCreated, result.Code)

	//unmarshal the response and assert
	response := new(models.Post)
	err := jsonapi.UnmarshalPayload(result.Body, response)
	require.NoError(t, err)

	require.Equal(t, post.ID, response.ID)
	require.Equal(t, post.Title, response.Title)
	require.Equal(t, post.ViewCount, response.ViewCount)
	require.WithinDuration(t, post.CreatedAt, response.CreatedAt, time.Second)
}

func TestGetPost(t *testing.T) {
	post := utils.FixturePost()

	in := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in, post); err != nil {
		log.Fatal(err)
	}

	postReq := httptest.NewRequest(http.MethodPost, "/posts", in)
	postReq.Header.Set(HeaderAccept, jsonapi.MediaType)

	router, _ := New()
	postResult := httptest.NewRecorder()
	router.CreatePost(postResult, postReq)

	id := post.ID
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/posts/%s", id), nil)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	//function under test
	result := httptest.NewRecorder()
	router.GetPost(result, req)

	require.Equal(t, http.StatusOK, result.Code)

	//unmarshal the response and assert
	response := new(models.Post)
	err := jsonapi.UnmarshalPayload(result.Body, response)
	require.NoError(t, err)

	require.Equal(t, post.ID, response.ID)
	require.Equal(t, post.Title, response.Title)
	require.Equal(t, post.ViewCount, response.ViewCount)
	require.WithinDuration(t, post.CreatedAt, response.CreatedAt, time.Second)
}
