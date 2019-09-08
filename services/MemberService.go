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

// MemberService logic service for member
type MemberService struct {
	MemberRepository interfaces.IMemberRepository
}

// Publish publish member
func (s *MemberService) Publish(payload models.MemberModel) (response models.APIResponseModel) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "Publish",
			constants.LogFieldPayload: payload,
		}

		member models.MemberModel
		err    error
	)

	// request validation
	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.OrgName, validation.Required, validation.Length(3, 100), validation.By(helpers.AlphaNum)),
		validation.Field(&payload.Username, validation.Required),
		validation.Field(&payload.AvatarURL, validation.Required),
		validation.Field(&payload.TotalFollower, validation.Required),
		validation.Field(&payload.TotalFollowing, validation.Required),
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
	member, err = s.MemberRepository.Publish(payload)
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

	logFields[constants.LogFieldResult] = member
	logFields[constants.LogFieldResponse] = response
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// GetAll get all member
func (s *MemberService) GetAll(payload models.MemberModel) (response models.APIResponseModel) {
	var (
		logFields = log.Fields{
			constants.LogFieldEvent:   "GetAll",
			constants.LogFieldPayload: payload,
		}
		members      []models.MemberModel
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

	members, err = s.MemberRepository.GetAll(payload)
	if err != nil {

		bodyResponse.Message = constants.MessageInternalError
		response.HTTPStatus = http.StatusInternalServerError
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Errorf(constants.LogMessageBasicError, err)
		return
	}

	// if not found
	if members == nil {
		bodyResponse.Message = constants.MessageMemberNotFound
		response.HTTPStatus = http.StatusNotFound
		response.BodyResponse = bodyResponse

		logFields[constants.LogFieldResponse] = response
		log.WithFields(logFields).Warnf(constants.LogMessageBasicWarning, bodyResponse)
		return
	}

	response.HTTPStatus = http.StatusOK
	response.BodyResponse = members

	logFields[constants.LogFieldResponse] = response
	log.WithFields(logFields).Info(constants.LogMessageBasicSuccess)
	return
}

// Delete delete member
func (s *MemberService) Delete(payload models.MemberModel) (response models.APIResponseModel) {
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

	deletedCount, err = s.MemberRepository.Delete(payload)
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
		bodyResponse.Message = constants.MessageMemberNotFound
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
