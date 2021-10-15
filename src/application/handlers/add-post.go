package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	usecases "github.com/ascartezini/go-discussion-forum/src/application/use-cases"
	"github.com/ascartezini/go-discussion-forum/src/domain/models"
)

type AddPostHandler struct {
	AddPostUseCase usecases.AddPostUseCase
}

func (h *AddPostHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			// w.Write(methodNotAllowedResponse)
			return
		}

		if r.Body == nil {
			log.Println("create requires a request body")
			w.WriteHeader(http.StatusBadRequest)
			// w.Write(badRequestResponse)
			return
		}

		var post models.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to create a post: %v\n", err)
			log.Println(msg)
			w.WriteHeader(http.StatusBadRequest)
			// w.Write(badRequestResponse)
			return
		}

		err = h.AddPostUseCase.AddPost(post)
		if err != nil {
			msg := fmt.Sprintf("an error occurred while trying to create a post: %v\n", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"message": %s}`, msg)))
			return
		}

		json.NewEncoder(w).Encode(post)
	}
}
