package interfaces

import "go-mongo/models"

// IMemberService interface
type IMemberService interface {
	Publish(payload models.MemberModel) (response models.APIResponseModel)
	GetAll(payload models.MemberModel) (response models.APIResponseModel)
	Delete(payload models.MemberModel) (response models.APIResponseModel)
}
