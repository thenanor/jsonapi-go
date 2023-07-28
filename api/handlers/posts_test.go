package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestGetPostWithoutIncluded(t *testing.T) {
	//prepare a post
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

	//start assertions
	require.Equal(t, http.StatusOK, result.Code)

	// duplicate the stream instead of copying io.Copy(body, result.Body)
	body := bytes.NewBuffer(nil)
	tee := io.TeeReader(result.Body, body)

	//assert we don't have "included" in the Body
	payload := new(jsonapi.OnePayload)
	err := json.NewDecoder(tee).Decode(payload)
	require.NoError(t, err)
	require.Nil(t, payload.Included)

	//unmarshal the response
	response := new(models.Post)
	err = jsonapi.UnmarshalPayload(body, response)
	require.NoError(t, err)

	require.Equal(t, post.ID, response.ID)
	require.Equal(t, post.Title, response.Title)
	require.Equal(t, post.ViewCount, response.ViewCount)
	require.WithinDuration(t, post.CreatedAt, response.CreatedAt, time.Second)
}

func TestGetPostWithIncluded(t *testing.T) {
	//prepare a post
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
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/posts/%s?include=author", id), nil)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	//function under test
	result := httptest.NewRecorder()
	router.GetPost(result, req)

	//start assertions
	require.Equal(t, http.StatusOK, result.Code)

	// duplicate the stream instead of copying io.Copy(body, result.Body)
	body := bytes.NewBuffer(nil)
	tee := io.TeeReader(result.Body, body)

	//assert we have "included" in the Body
	payload := new(jsonapi.OnePayload)
	err := json.NewDecoder(tee).Decode(payload)
	require.NoError(t, err)
	require.NotNil(t, payload.Included)

	//unmarshal the response
	response := new(models.Post)
	err = jsonapi.UnmarshalPayload(body, response)
	require.NoError(t, err)

	require.Equal(t, post.ID, response.ID)
	require.Equal(t, post.Title, response.Title)
	require.Equal(t, post.ViewCount, response.ViewCount)
	require.WithinDuration(t, post.CreatedAt, response.CreatedAt, time.Second)
}

func TestGetPosts(t *testing.T) {
	//prepare posts
	post1 := utils.FixturePost()
	in1 := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in1, post1); err != nil {
		log.Fatal(err)
	}

	post1Req := httptest.NewRequest(http.MethodPost, "/posts", in1)
	post1Req.Header.Set(HeaderAccept, jsonapi.MediaType)

	post2 := utils.FixturePost()
	in2 := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in2, post2); err != nil {
		log.Fatal(err)
	}

	post2Req := httptest.NewRequest(http.MethodPost, "/posts", in2)
	post2Req.Header.Set(HeaderAccept, jsonapi.MediaType)

	post3 := utils.FixturePost()
	in3 := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in3, post3); err != nil {
		log.Fatal(err)
	}

	post3Req := httptest.NewRequest(http.MethodPost, "/posts", in3)
	post3Req.Header.Set(HeaderAccept, jsonapi.MediaType)

	router, _ := New()
	post1Result := httptest.NewRecorder()
	router.CreatePost(post1Result, post1Req)
	post2Result := httptest.NewRecorder()
	router.CreatePost(post2Result, post2Req)
	post3Result := httptest.NewRecorder()
	router.CreatePost(post3Result, post3Req)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	//function under test
	result := httptest.NewRecorder()
	router.GetPosts(result, req)

	//start assertions
	require.Equal(t, http.StatusOK, result.Code)

	//unmarshal the response
	models, err := jsonapi.UnmarshalManyPayload(result.Body, reflect.TypeOf(new(models.Post)))
	require.NoError(t, err)
	require.NotNil(t, models)
	require.Equal(t, 3, len(models))
}
