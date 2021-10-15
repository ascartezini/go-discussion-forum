package usecases

import (
	"github.com/ascartezini/go-discussion-forum/src/domain/models"
	"github.com/ascartezini/go-discussion-forum/src/domain/repositories"
)

type AddPostUseCase struct {
	PostRepository repositories.PostRepository
}

func (usecase *AddPostUseCase) AddPost(post models.Post) error {
	err := usecase.PostRepository.Add(post)

	if err != nil {
		println(err)
	}

	return nil
}
