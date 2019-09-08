package interfaces

import (
	"go-mongo/models"
)

// ICommentRepository interface
type ICommentRepository interface {
	Publish(payload models.CommentModel) (result models.CommentModel, err error)
	GetAll(payload models.CommentModel) (comments []models.CommentModel, err error)
	Delete(payload models.CommentModel) (deletedCount int, err error)
}
