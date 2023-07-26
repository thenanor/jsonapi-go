package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/thenanor/jsonapi-go/api/handlers"
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
