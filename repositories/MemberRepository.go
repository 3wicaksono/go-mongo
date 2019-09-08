package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"go-mongo/constants"
	"go-mongo/infrastructures"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

// MemberRepository struct
type MemberRepository struct {
	MongoConnect infrastructures.MongoConnect
}

// Publish publish member
func (r *MemberRepository) Publish(payload models.MemberModel) (result models.MemberModel, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Publish",
			constants.LogFieldPayload: payload,
		}
		collection = r.MongoConnect.Database.Collection(constants.MongoCollectionMembers)
		ctx, _     = context.WithTimeout(context.Background(), constants.MongoMaxTimeout*time.Second)
	)

	payload.IsDeleted = 0
	insert, err := collection.InsertOne(ctx, payload)
	if err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	// get object id string
	oid, _ := insert.InsertedID.(primitive.ObjectID)
	result.ID = oid
	result.OrgName = payload.OrgName
	result.Username = payload.Username
	result.AvatarURL = payload.AvatarURL
	result.TotalFollower = payload.TotalFollower
	result.TotalFollowing = payload.TotalFollowing
	logFields[constants.LogFieldResult] = result
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// GetAll get all member
func (r *MemberRepository) GetAll(payload models.MemberModel) (members []models.MemberModel, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "GetAll",
			constants.LogFieldPayload: payload,
		}
		collection = r.MongoConnect.Database.Collection(constants.MongoCollectionMembers)
		ctx, _     = context.WithTimeout(context.Background(), constants.MongoMaxTimeout*time.Second)
		cursor     *mongo.Cursor
	)

	options := options.Find()

	// Sort by `total_follower` field descending
	options.SetSort(bson.D{{"total_follower", -1}})

	if payload.OrgName != "" {
		cursor, err = collection.Find(ctx, models.CommentModel{OrgName: payload.OrgName, IsDeleted: 0}, options)
	} else {
		cursor, err = collection.Find(ctx, models.CommentModel{IsDeleted: 0}, options)
	}

	if err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var member models.MemberModel
		cursor.Decode(&member)
		members = append(members, member)
	}

	if err = cursor.Err(); err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	logFields[constants.LogFieldResult] = members
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// Delete soft delete member
func (r *MemberRepository) Delete(payload models.MemberModel) (deletedCount int, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Delete",
			constants.LogFieldPayload: payload,
		}
	)
	//idMongo, _ := primitive.ObjectIDFromHex(id)
	collection := r.MongoConnect.Database.Collection(constants.MongoCollectionMembers)
	ctx, _ := context.WithTimeout(context.Background(), constants.MongoMaxTimeout*time.Second)

	update := bson.D{{"$set",
		bson.D{
			{"is_deleted", 1},
		},
	}}

	result, err := collection.UpdateMany(ctx, models.CommentModel{OrgName: payload.OrgName}, update)
	if err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	deletedCount = int(result.ModifiedCount)

	logFields[constants.LogFieldResult] = result
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}
