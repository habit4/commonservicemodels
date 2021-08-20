package commonservicemodels

import (
	"github.com/habit4/context"
	"github.com/habit4/errors"
	"reflect"
)

//ServiceResponse  common service response structure
type ServiceResponse struct {
	Messages   []Message
	Data       []interface{}
	Pagination PaginationData
	Success    bool
	StatusCode string
}

//JSON convert to json
func (s *ServiceResponse) JSON() (json string) {
	if len(s.Messages) == 0 {
		s.Messages = make([]Message, 0)
	}
	if len(s.Data) == 0 {
		s.Data = make([]interface{}, 0)
	}

	if len(s.Messages) == 0 {
		s.Success = true
	}

	if s.Success == true {
		s.StatusCode = "200"
	}

	if s.Success == false {
		s.StatusCode = "500"
	}

	json, _ = ToJSON(s)

	return
}

//AddError add error object
func (s *ServiceResponse) AddError(err error) *ServiceResponse {
	if err != nil {
		s.Success = false
		customError, ok := err.(errors.Error)
		if ok {
			s.Messages = MessagesFromError(customError)
		} else {

			s.Messages = append(s.Messages, Message{
				Message: err.Error(),
				Type:    "error",
			})
		}

	}
	return s
}

//AddData add data
func (s *ServiceResponse) AddData(data interface{}) *ServiceResponse {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			s.Data = append(s.Data, val.Index(i).Interface())
		}
	} else {
		s.Data = append(s.Data, data)
	}

	return s

}

//SetData set dta
func (s *ServiceResponse) SetData(data []interface{}) *ServiceResponse {
	s.Data = append(s.Data, data...)
	return s
}

//AddInfoMessage add info message
func (s *ServiceResponse) AddInfoMessage(message string) *ServiceResponse {
	s.Success = true
	messageStruct := Message{}
	s.Messages = append(s.Messages, *messageStruct.Info(message))
	return s
}

func (s *ServiceResponse) AddErrorMessage(message string) *ServiceResponse {
	s.Success = false
	messageStruct := Message{}
	s.Messages = append(s.Messages, *messageStruct.ErrorString(message))
	return s
}

//PaginationData data for pagination
type PaginationData struct {
	TotalCount int
	Skip       int
	Limit      int
}

//ServiceRequestData service request data
type ServiceRequestData struct {
	ExecutionContext context.ExecutionContextDetails
	PayloadJSON      string
	Payload          interface{}
}

//Validate validate request data
func (s ServiceRequestData) Validate() (response ServiceResponse) {
	response.Success = true
	/**err := s.ExecutionContext.Validate()
	if err != nil {
		response.AddError(err)
	}**/ //TODO add validations
	return
}

//AddUnauthorizedError unauthorized
func (s *ServiceResponse) AddUnauthorizedError() {
	s.Success = false
	s.StatusCode = "401"
	s.AddError(errors.Error{
		Message: "Not Authorized",
		Code:    "NOT_AUTHORIZED",
	})
}
