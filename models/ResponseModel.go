package models

// ResponseModel - struct of Response
type ResponseModel struct {
	Message string `json:"message"`
}

// APIResponseModel struct of API Response
type APIResponseModel struct {
	HTTPStatus   int
	BodyResponse interface{}
}

// APIMessage struct of API Message
type APIMessage struct {
	Message string `json:"message"`
}
