package alexa

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//----------------------------------------------------------------------------------------------------------------------
type RequestType string

const (
	LaunchRequestType       RequestType = "LaunchRequest"
	IntentRequestType       RequestType = "IntentRequest"
	SessionEndedRequestType RequestType = "SessionEndedRequest"
)

//----------------------------------------------------------------------------------------------------------------------
// Request interface
type Request interface {
	RequestType() RequestType
}

//----------------------------------------------------------------------------------------------------------------------
func ParseRequest(b json.RawMessage) (Request, error) {
	var request Request
	var rawReq struct {
		Type RequestType `json:"type"`
	}
	err := json.Unmarshal(b, &rawReq)
	if err != nil {
		return nil, err
	}
	switch rawReq.Type {
	case LaunchRequestType:
		request = &LaunchRequest{}
	case IntentRequestType:
		request = &IntentRequest{}
	case SessionEndedRequestType:
		request = &SessionEndRequest{}
	default:
		return nil, errors.New(fmt.Sprintf("unknown request type: %s", rawReq.Type))
	}

	err = json.Unmarshal(b, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

//----------------------------------------------------------------------------------------------------------------------
// LaunchRequest type
type LaunchRequest struct {
	Type      RequestType `json:"type"`
	RequestId string      `json:"requestId"`
	TimeStamp time.Time   `json:"timestamp"`
	Locale    string      `json:"locale"`
}

// LaunchRequest implements Request interface
func (req LaunchRequest) RequestType() RequestType { return req.Type }

//----------------------------------------------------------------------------------------------------------------------
// IntentRequest type
type IntentRequest struct {
	Type        RequestType     `json:"type"`
	RequestId   string          `json:"requestId"`
	TimeStamp   time.Time       `json:"timestamp"`
	Locale      string          `json:"locale"`
	DialogState DialogStateType `json:"dialogState"`
	Intent      Intent          `json:"intent"`
}

// IntentRequest implements Request interface
func (req IntentRequest) RequestType() RequestType { return req.Type }

type DialogStateType string

const (
	STARTED    DialogStateType = "STARTED"
	INPROGRESS DialogStateType = "IN_PROGRESS"
	COMPLETED  DialogStateType = "COMPLETED"
)

type ConfirmationStatusType string

const (
	NONE      ConfirmationStatusType = "NONE"
	CONFIRMED ConfirmationStatusType = "CONFIRMED"
	DENIED    ConfirmationStatusType = "DENIED"
)

type Intent struct {
	Name               string                 `json:"name"`
	ConfirmationStatus ConfirmationStatusType `json:"confirmationStatus"`
	Slots              map[string]Slot        `json:"slots"`
}

type Slot struct {
	Name               string                 `json:"name"`
	Value              string                 `json:"value"`
	ConfirmationStatus ConfirmationStatusType `json:"confirmationStatus"`
	// Todo: resolutions
	// Resolutions interface{} `json:"resolutions"`
}

//----------------------------------------------------------------------------------------------------------------------
// SessionEndRequest type
type SessionEndRequest struct {
	Type      RequestType            `json:"type"`
	RequestId string                 `json:"requestId"`
	TimeStamp time.Time              `json:"timestamp"`
	Locale    string                 `json:"locale"`
	Reason    SessionEndReqReason    `json:"reason"`
	Error     SessionEndRequestError `json:"error"`
}

// SessionEndRequest implements Request interface
func (req SessionEndRequest) RequestType() RequestType { return req.Type }

type SessionEndReqReason string

const (
	USER_INITIATED         SessionEndReqReason = "USER_INITIATED"
	ERROR                  SessionEndReqReason = "ERROR"
	EXCEEDED_MAX_REPROMPTS SessionEndReqReason = "EXCEEDED_MAX_REPROMPTS"
)

type SessionEndRequestErrorType string

const (
	INVALID_RESPONSE           SessionEndRequestErrorType = "INVALID_RESPONSE"
	DEVICE_COMMUNICATION_ERROR SessionEndRequestErrorType = "DEVICE_COMMUNICATION_ERROR"
	INTERNAL_SERVICE_ERROR     SessionEndRequestErrorType = "INTERNAL_SERVICE_ERROR"
	ENDPOINT_TIMEOUT           SessionEndRequestErrorType = "ENDPOINT_TIMEOUT"
)

type SessionEndRequestError struct {
	Type    SessionEndRequestErrorType `json:"type"`
	Message string                     `json:"message"`
}

//----------------------------------------------------------------------------------------------------------------------
