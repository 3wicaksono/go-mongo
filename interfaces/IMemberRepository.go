package interfaces

import "go-mongo/models"

// IMemberRepository interface
type IMemberRepository interface {
	Publish(payload models.MemberModel) (result models.MemberModel, err error)
	GetAll(payload models.MemberModel) (members []models.MemberModel, err error)
	Delete(payload models.MemberModel) (deletedCount int, err error)
}
