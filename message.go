package commonservicemodels

import "github.com/habit4/errors"

//Message structure
type Message struct {
	Type    string
	Message string
	Code    string
}

func MessagesFromError(err errors.Error) (messages []Message) {
	if len(err.Messages) > 0 {
		for _, message := range err.Messages {
			message := Message{
				Message: message,
				Type:    "error",
			}
			messages = append(messages, message)
		}
	} else {
		message := Message{
			Message: err.Message,
			Type:    "error",
		}
		messages = append(messages, message)
	}

	return messages
}

func (m *Message) Error(err error) *Message {
	m.Type = "error"
	m.Message = err.Error()
	return m
}

func (m *Message) ErrorString(err string) *Message {
	m.Type = "error"
	m.Message = err
	return m
}

//Info info
func (m *Message) Info(info string) *Message {
	m.Type = "info"
	m.Message = info
	return m
}

//Warn message
func (m *Message) Warn(warn string) *Message {
	m.Type = "warn"
	m.Message = warn
	return m
}

//Fatal error message
func (m *Message) Fatal(fatal string) *Message {
	m.Type = "error"
	m.Message = fatal
	return m
}
