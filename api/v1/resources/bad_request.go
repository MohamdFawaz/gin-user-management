package resources

import "net/http"

type BadRequest struct {
	Status int
	Message string
	Data map[string]interface{}
}

func BadRequestResource(message string) IResource  {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := BadRequest{Status: http.StatusBadRequest, Message: message, Data: data}
	return &resource
}

func (badRequest BadRequest) GetStatus() int {
	return badRequest.Status
}

func (badRequest BadRequest) GetData() map[string]interface{} {
	return badRequest.Data
}