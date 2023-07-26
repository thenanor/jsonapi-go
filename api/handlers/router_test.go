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

func TestServeHTTP(t *testing.T) {
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

	result := httptest.NewRecorder()
	router, err := New()
	require.NoError(t, err)

	//function under test
	router.ServeHTTP(result, req)

	//print the response payload in case of error
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, result.Body.Bytes(), "", "  ")
	fmt.Println(prettyJSON.String())

	require.Equal(t, http.StatusCreated, result.Code)
}
