package alexa

//----------------------------------------------------------------------------------------------------------------------
// Response type
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#response-parameters
type ResponseBody struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
	Response          Response               `json:"response"`
}

//----------------------------------------------------------------------------------------------------------------------
type Response interface {
	AlexaResponse()
}

//----------------------------------------------------------------------------------------------------------------------
func Ask(msg string) *BasicResponse {
	return &BasicResponse{
		OutputSpeech: OutputSpeech{
			Type: PLAIN_TEXT_TYPE,
			Text: msg,
		},
		Card:             nil,
		Reprompt:         nil,
		ShouldEndSession: false,
		Directives:       nil,
	}
}

func Tell(msg string) *BasicResponse {
	return &BasicResponse{
		OutputSpeech: OutputSpeech{
			Type: PLAIN_TEXT_TYPE,
			Text: msg,
		},
		Card:             nil,
		Reprompt:         nil,
		ShouldEndSession: true,
		Directives:       nil,
	}
}

//----------------------------------------------------------------------------------------------------------------------
// ResponseBody type
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#response-object
type BasicResponse struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Card             Card         `json:"card,omitempty"`
	Reprompt         *Reprompt    `json:"reprompt,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
	//todo: implement directives struct
	Directives interface{} `json:"directives,omitempty"`
}

// BasicResponse implements alexa.Response
func (BasicResponse) AlexaResponse() {}

//----------------------------------------------------------------------------------------------------------------------
// OutputSpeech type
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#outputspeech-object
type OutputSpeechType string

const (
	PLAIN_TEXT_TYPE OutputSpeechType = "PlainText"
	SSML_TYPE       OutputSpeechType = "SSML"
)

type OutputSpeech struct {
	Type         OutputSpeechType `json:"type"`
	Text         string           `json:"text,omitempty"`
	SSML         string           `json:"ssml,omitempty"`
	PlayBehavior PlayBehaviorType `json:"playBehavior,omitempty"`
}

type PlayBehaviorType string

const (
	ENQUEUE          PlayBehaviorType = "ENQUEUE"
	REPLACE_ALL      PlayBehaviorType = "REPLACE_ALL"
	REPLACE_ENQUEUED PlayBehaviorType = "REPLACE_ENQUEUED"
)

//----------------------------------------------------------------------------------------------------------------------
// Reprompt type
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#reprompt-object
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}
