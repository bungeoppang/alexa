package alexa

import (
	"encoding/json"
	"fmt"
	"testing"
)

var unmarshalErr = "Unmarshal Error"

func getABNotEqualErrMsg(a, b interface{}) string {
	return fmt.Sprintf("Not Equal: %5s != %5s", a, b)
}

func Test_UnmarshalLaunchRequest(t *testing.T) {
	var request LaunchRequest
	rawData := []byte(`{
			"type": "LaunchRequest",
			"requestId": "amzn1.echo-api.request.1",
			"timestamp": "2016-10-27T18:21:44Z",
			"locale": "en-US"}`,
	)
	err := json.Unmarshal(rawData, &request)
	if err != nil {
		t.Error(err.Error())
	}

	if request.Type != LaunchRequestType {
		t.Error(getABNotEqualErrMsg(request.Type, LaunchRequestType))
	}

	if request.RequestId != "amzn1.echo-api.request.1" {
		t.Error(getABNotEqualErrMsg(request.RequestId, "amzn1.echo-api.request.1"))
	}

	if request.Locale != "en-US" {
		t.Error(getABNotEqualErrMsg(request.Locale, "en-US"))
	}
}

func Test_UnmarshalIntentRequest(t *testing.T) {
	var request IntentRequest
	rawData := []byte(`{
  		"type": "IntentRequest",
  		"requestId": "testId",
  		"timestamp": "2016-10-27T18:21:44Z",
  		"dialogState": "STARTED",
  		"locale": "ko-KR",
  		"intent": {
    		"name": "test-intent",
    		"confirmationStatus": "NONE",
    		"slots": {
      			"test-slot": {
        			"name": "test-slot",
        			"value": "test-slot-value",
        			"confirmationStatus": "NONE"
      			}
    		}
  		}
	}`)
	err := json.Unmarshal(rawData, &request)
	if err != nil {
		t.Error(err.Error())
	}

	if request.Type != IntentRequestType {
		t.Error(getABNotEqualErrMsg(request.Type, IntentRequestType))
	}

	if request.RequestId != "testId" {
		t.Error(getABNotEqualErrMsg(request.RequestId, "testId"))
	}

	if request.Locale != "ko-KR" {
		t.Error(getABNotEqualErrMsg(request.Locale, "en-US"))
	}

	if request.Intent.Name != "test-intent" {
		t.Error(getABNotEqualErrMsg(request.Intent.Name, "test-intent"))
	}

	if request.Intent.ConfirmationStatus != NONE {
		t.Error(getABNotEqualErrMsg(request.Intent.ConfirmationStatus, NONE))
	}

	slots := request.Intent.Slots
	if len(slots) != 1 {
		t.Error(unmarshalErr)
	}

	slot, exist := slots["test-slot"]
	if !exist {
		t.Error(unmarshalErr)
	}

	if slot.ConfirmationStatus != NONE || slot.Name != "test-slot" || slot.Value != "test-slot-value" {
		t.Error(unmarshalErr)
	}
}

func Test_UnmarshalSessionEndRequest(t *testing.T) {
	var request SessionEndRequest
	rawData := []byte(`{
		"type": "SessionEndedRequest",
  		"requestId": "test-req-id",
  		"timestamp": "2016-10-27T18:21:44Z",
  		"reason": "USER_INITIATED",
  		"locale": "ja-JP",
  		"error": {
    		"type": "DEVICE_COMMUNICATION_ERROR",
    		"message": "test message"
  		}
	}`)

	err := json.Unmarshal(rawData, &request)
	if err != nil {
		t.Error(err.Error())
	}

	if request.Type != SessionEndedRequestType {
		t.Error(unmarshalErr)
	}

	if request.RequestId != "test-req-id" {
		t.Error(unmarshalErr)
	}

	if request.Locale != "ja-JP" {
		t.Error(unmarshalErr)
	}

	if request.Reason != USER_INITIATED {
		t.Error(unmarshalErr)
	}

	if request.Error.Type != DEVICE_COMMUNICATION_ERROR {
		t.Error(unmarshalErr)
	}

	if request.Error.Message != "test message" {
		t.Error(unmarshalErr)
	}
}
