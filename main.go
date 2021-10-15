package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ascartezini/go-discussion-forum/src/application/handlers"
	usecases "github.com/ascartezini/go-discussion-forum/src/application/use-cases"
	datasources "github.com/ascartezini/go-discussion-forum/src/infrastructure/data-sources"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"service is up and running"}`))
	})

	db := datasources.DynamoDBPostRepository{}
	usecase := usecases.AddPostUseCase{PostRepository: db}
	h := handlers.AddPostHandler{AddPostUseCase: usecase}

	router.HandleFunc("/posts", h.Handle())

	srv := http.Server{
		Addr:         ":8080",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		Handler:      router,
	}

	fmt.Println("http server listening on localhost:8080")
	log.Fatal(srv.ListenAndServe())

}
