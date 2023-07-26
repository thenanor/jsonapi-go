package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/require"
	"github.com/thenanor/jsonapi-go/utils"
)

func TestServeHTTPForPost(t *testing.T) {
	post := utils.FixturePost()

	in := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in, post); err != nil {
		log.Fatal(err)
	}

	//print the request payload in case of error
	var prettyJSONRequest bytes.Buffer
	_ = json.Indent(&prettyJSONRequest, in.Bytes(), "", "  ")
	fmt.Println(prettyJSONRequest.String())

	//prepare the request to be sent to the handler
	req := httptest.NewRequest(http.MethodPost, "/posts", in)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	router, err := New()
	require.NoError(t, err)

	result := httptest.NewRecorder()

	//function under test
	router.ServeHTTP(result, req)

	//print the response payload in case of error
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, result.Body.Bytes(), "", "  ")
	fmt.Println(prettyJSON.String())

	require.Equal(t, http.StatusCreated, result.Code)
}

func TestServeHTTPForGet(t *testing.T) {
	post := utils.FixturePost()

	in := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in, post); err != nil {
		log.Fatal(err)
	}

	postReq := httptest.NewRequest(http.MethodPost, "/posts", in)
	postReq.Header.Set(HeaderAccept, jsonapi.MediaType)

	router, err := New()
	require.NoError(t, err)

	postResult := httptest.NewRecorder()
	router.ServeHTTP(postResult, postReq)
	require.Equal(t, http.StatusCreated, postResult.Code)

	id := post.ID
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/posts/%s", id), nil)
	req.Header.Set(HeaderAccept, jsonapi.MediaType)

	//function under test
	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)
	require.Equal(t, http.StatusOK, result.Code)
}
