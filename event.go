package alexa

import "encoding/json"

//----------------------------------------------------------------------------------------------------------------------
type Event struct {
	Version string       `json:"version"`
	Request Request      `json:"request"`
	Context EventContext `json:"context"`
	Session Session      `json:"session"`
}

func (event *Event) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Version string          `json:"version"`
		Request json.RawMessage `json:"request"`
		Context EventContext    `json:"context"`
		Session Session         `json:"session"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	parsedReq, err := ParseRequest(tmp.Request)
	if err != nil {
		return err
	}

	event.Version = tmp.Version
	event.Context = tmp.Context
	event.Request = parsedReq
	event.Session = tmp.Session

	return nil
}

//----------------------------------------------------------------------------------------------------------------------
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#session-object
type Session struct {
	New         bool                   `json:"new"`
	SessionId   string                 `json:"sessionId"`
	Attributes  map[string]interface{} `json:"attributes"`
	Application ApplicationObject      `json:"application"`
	User        User                   `json:"user"`
}

//----------------------------------------------------------------------------------------------------------------------
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#context-object
type EventContext struct {
	// todo: APL
	// APL reference: https://developer.amazon.com/en-US/docs/alexa/alexa-presentation-language/understand-apl.html
	AudioPlayer AudioPlayer `json:"AudioPlayer,omitempty"`
	System      System      `json:"System"`
}

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#system-object
type System struct {
	ApiAccessToken string            `json:"apiAccessToken"`
	ApiEndPoint    string            `json:"apiEndPoint"`
	Application    ApplicationObject `json:"application"`
	Device         Device            `json:"device"`
	Person         Person            `json:"person"`
	User           User              `json:"user"`
}

type ApplicationObject struct {
	ApplicationId string `json:"applicationId"`
}

type Device struct {
	DeviceId            string      `json:"deviceId"`
	SupportedInterfaces interface{} `json:"supportedInterfaces"`
}

type Person struct {
	PersonId    string `json:"personId"`
	AccessToken string `json:"accessToken,omitempty"`
}

type User struct {
	UserId      string `json:"userId"`
	AccessToken string `json:"accessToken,omitempty"`
}

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#audioplayer-object
type AudioPlayerActivityType string

const (
	IDLE            AudioPlayerActivityType = "IDLE"
	PAUSED          AudioPlayerActivityType = "PAUSED"
	PLAYING         AudioPlayerActivityType = "PLAYING"
	BUFFER_UNDERRUN AudioPlayerActivityType = "BUFFER_UNDERRUN"
	FINISHED        AudioPlayerActivityType = "FINISHED"
	STOPPED         AudioPlayerActivityType = "STOPPED"
)

type AudioPlayer struct {
	Token                string                  `json:"token"`
	OffsetInMilliseconds float64                 `json:"offsetInMilliseconds"`
	PlayerActivity       AudioPlayerActivityType `json:"playerActivity"`
}
