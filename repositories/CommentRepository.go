package repositories

import (
	"context"
	"time"

	"go-mongo/constants"
	"go-mongo/infrastructures"
	"go-mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

// CommentRepository struct
type CommentRepository struct {
	MongoConnect infrastructures.MongoConnect
}

// Publish publish comment
func (r *CommentRepository) Publish(payload models.CommentModel) (result models.CommentModel, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Publish",
			constants.LogFieldPayload: payload,
		}
		collection = r.MongoConnect.Database.Collection(constants.MongoCollectionComments)
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
	result.Comment = payload.Comment
	logFields[constants.LogFieldResult] = result
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// GetAll get all comment
func (r *CommentRepository) GetAll(payload models.CommentModel) (comments []models.CommentModel, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "GetAll",
			constants.LogFieldPayload: payload,
		}
		collection = r.MongoConnect.Database.Collection(constants.MongoCollectionComments)
		ctx, _     = context.WithTimeout(context.Background(), constants.MongoMaxTimeout*time.Second)
		cursor     *mongo.Cursor
	)

	if payload.OrgName != "" {
		cursor, err = collection.Find(ctx, models.CommentModel{OrgName: payload.OrgName, IsDeleted: 0})
	} else {
		cursor, err = collection.Find(ctx, models.CommentModel{IsDeleted: 0})
	}

	if err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.CommentModel
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}

	if err = cursor.Err(); err != nil {
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	logFields[constants.LogFieldResult] = comments
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// Delete soft delete comment
func (r *CommentRepository) Delete(payload models.CommentModel) (deletedCount int, err error) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Delete",
			constants.LogFieldPayload: payload,
		}
	)
	//idMongo, _ := primitive.ObjectIDFromHex(id)
	collection := r.MongoConnect.Database.Collection(constants.MongoCollectionComments)
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
