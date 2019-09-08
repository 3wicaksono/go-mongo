package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentModel model for comment
type CommentModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrgName   string             `json:"org_name,omitempty" bson:"org_name,omitempty"`
	Comment   string             `json:"comment,omitempty" bson:"comment,omitempty"`
	IsDeleted int                `json:"is_deleted,omitempty" bson:"is_deleted"`
}
