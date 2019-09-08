package interfaces

import "go-mongo/models"

// ICommentService interface
type ICommentService interface {
	Publish(payload models.CommentModel) (response models.APIResponseModel)
	GetAll(payload models.CommentModel) (response models.APIResponseModel)
	Delete(payload models.CommentModel) (response models.APIResponseModel)
}
