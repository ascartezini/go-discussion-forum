package datasources

import (
	"github.com/ascartezini/go-discussion-forum/src/domain/models"
)

type DynamoDBPostRepository struct{}

func (db DynamoDBPostRepository) Add(p models.Post) error {
	return nil
}

func (db DynamoDBPostRepository) Delete(id int) error {
	return nil
}
