package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberModel model for member
type MemberModel struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrgName        string             `json:"org_name,omitempty" bson:"org_name,omitempty"`
	Username       string             `json:"username,omitempty" bson:"username,omitempty"`
	AvatarURL      string             `json:"avatar_url,omitempty" bson:"avatar_url,omitempty"`
	TotalFollower  int                `json:"total_follower,omitempty" bson:"total_follower,omitempty"`
	TotalFollowing int                `json:"total_following,omitempty" bson:"total_following,omitempty"`
	IsDeleted      int                `json:"is_deleted,omitempty" bson:"is_deleted"`
}
