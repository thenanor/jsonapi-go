package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"time"

	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"github.com/thenanor/jsonapi-go/api/handlers"
	"github.com/thenanor/jsonapi-go/api/models"
)

func main() {
	router, err := handlers.New()
	if err != nil {
		log.Fatalf("unable to create api router %v", err)
	}
	http.HandleFunc("/posts", router.ServeHTTP)

	sv := http.Server{
		Addr:         ":9090",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		err := sv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("============ POST a post ===========")
	post := &models.Post{
		ID:        uuid.NewString(),
		Title:     "My First Post",
		CreatedAt: time.Now(),
		// Comments:  []*models.Comment{},
	}

	in := bytes.NewBuffer(nil)
	if err := jsonapi.MarshalOnePayloadEmbedded(in, post); err != nil {
		log.Fatal(err)
	}
	fmt.Println("thepayload:", in)
	fmt.Println(post)

	req, _ := http.NewRequest(http.MethodPost, "/posts", in)
	req.Header.Set(handlers.HeaderAccept, jsonapi.MediaType)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)
	fmt.Println("============ stop POST a post ===========")

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, w.Body)

	fmt.Println("============ jsonapi response from create ===========")
	fmt.Println(buf.String())
	fmt.Println("============== end raw jsonapi response =============")

	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, os.Kill)
	<-sigs
	log.Println("Received terminate signal - starting graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := sv.Shutdown(ctx); err != nil {
		log.Fatalf("received an error %v", err)
	}
	log.Println("shutdown complete")
	os.Exit(0)
}
