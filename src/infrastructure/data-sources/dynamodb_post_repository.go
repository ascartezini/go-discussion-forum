package datasources

import (
	"github.com/ascartezini/go-discussion-forum/src/domain/models"
)

type DynamoDBPostRepository struct{}

func (db DynamoDBPostRepository) Add(p models.Post) (string, error) {
	return "a53d9629-44c6-4074-8ff9-8f7b73928eb5", nil
}

func (db DynamoDBPostRepository) Delete(id int) error {
	return nil
}
