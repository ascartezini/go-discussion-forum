package repositories

import (
	"github.com/ascartezini/go-discussion-forum/src/domain/models"
)

type PostRepository interface {
	Add(p models.Post) (string, error)
	Delete(id int) error
}
