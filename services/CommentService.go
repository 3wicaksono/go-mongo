package services

import (
	"go-mongo/helpers"
	"net/http"

	"go-mongo/constants"
	"go-mongo/interfaces"
	"go-mongo/models"

	validation "github.com/go-ozzo/ozzo-validation"
	log "github.com/sirupsen/logrus"
)

// CommentService logic service for comment
type CommentService struct {
	CommentRepository interfaces.ICommentRepository
}

// Publish publish comment
func (s *CommentService) Publish(payload models.CommentModel) (response models.APIResponseModel) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Publish",
			constants.LogFieldPayload: payload,
		}

		comment models.CommentModel
		err     error
	)

	// request validation
	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.OrgName, validation.Required, validation.Length(3, 100), validation.By(helpers.AlphaNum)),
		validation.Field(&payload.Comment, validation.Required),
	)

	// if validate is not passed
	if err != nil {
		response.HTTPStatus = http.StatusBadRequest
		response.BodyResponse = models.APIMessage{
			Message: err.Error(),
		}

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, err)
		return
	}

	// publish to DB
	comment, err = s.CommentRepository.Publish(payload)
	if err != nil {
		response.HTTPStatus = http.StatusInternalServerError
		response.BodyResponse = models.APIMessage{
			Message: constants.MessageInternalError,
		}

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	response.HTTPStatus = http.StatusOK
	response.BodyResponse = models.APIMessage{
		Message: constants.MessageSuccess,
	}

	logFields[constants.LogFieldResult] = comment
	logFields[constants.LogFieldResponse] = response
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// GetAll get all comment
func (s *CommentService) GetAll(payload models.CommentModel) (response models.APIResponseModel) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "GetAll",
			constants.LogFieldPayload: payload,
		}
		comments     []models.CommentModel
		err          error
		bodyResponse models.APIMessage
	)

	if payload.OrgName != "" {
		// request validation
		err = validation.ValidateStruct(&payload,
			validation.Field(&payload.OrgName, validation.Required, validation.Length(3, 100), validation.By(helpers.AlphaNum)),
		)

		// if validate is not passed
		if err != nil {
			bodyResponse.Message = err.Error()

			response.HTTPStatus = http.StatusBadRequest
			response.BodyResponse = bodyResponse

			logFields[constants.LogFieldResponse] = response
			log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, err)
			return
		}
	}

	comments, err = s.CommentRepository.GetAll(payload)
	if err != nil {

		bodyResponse.Message = constants.MessageInternalError
		response.HTTPStatus = http.StatusInternalServerError
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	// if not found
	if comments == nil {
		bodyResponse.Message = constants.MessageCommentNotFound
		response.HTTPStatus = http.StatusNotFound
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, bodyResponse)
		return
	}

	response.HTTPStatus = http.StatusOK
	response.BodyResponse = comments

	logFields[constants.LogFieldResponse] = response
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// Delete delete comment
func (s *CommentService) Delete(payload models.CommentModel) (response models.APIResponseModel) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Delete",
			constants.LogFieldPayload: payload,
		}
		bodyResponse models.APIMessage
		deletedCount int
		err          error
	)

	// request validation
	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.OrgName, validation.Required, validation.Length(3, 100), validation.By(helpers.AlphaNum)),
	)

	// if validate is not passed
	if err != nil {
		response.HTTPStatus = http.StatusBadRequest
		response.BodyResponse = models.APIMessage{
			Message: err.Error(),
		}

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, err)
		return
	}

	deletedCount, err = s.CommentRepository.Delete(payload)
	if err != nil {

		bodyResponse.Message = constants.MessageInternalError
		response.HTTPStatus = http.StatusInternalServerError
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	// if not found
	if deletedCount == 0 {
		bodyResponse.Message = constants.MessageCommentNotFound
		response.HTTPStatus = http.StatusNotFound
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, bodyResponse)
		return
	}

	bodyResponse.Message = constants.MessageSuccess
	response.HTTPStatus = http.StatusOK
	response.BodyResponse = bodyResponse

	logFields[constants.LogFieldResponse] = response
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}
