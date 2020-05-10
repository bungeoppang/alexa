package alexa

import (
	"fmt"
)

//----------------------------------------------------------------------------------------------------------------------
func Skill(appId string) *skill {
	return &skill{
		AppId:                appId,
		intentRequestHandler: make(map[string]IntentRequestHandleFunc),
	}
}

//----------------------------------------------------------------------------------------------------------------------
type skill struct {
	AppId                    string
	defaultHandler           DefaultHandleFunc
	launchRequestHandler     LaunchRequestHandleFunc
	intentRequestHandler     map[string]IntentRequestHandleFunc
	sessionEndRequestHandler SessionEndRequestHandleFunc
}

//----------------------------------------------------------------------------------------------------------------------
func (s *skill) HandleDefaultFunc(f DefaultHandleFunc) {
	s.defaultHandler = f
}

func (s *skill) HandleLaunchRequestFunc(f LaunchRequestHandleFunc) {
	s.launchRequestHandler = f
}

func (s *skill) HandleIntentRequestFunc(intent string, f IntentRequestHandleFunc) {
	//todo: handle intent with regex pattern
	s.intentRequestHandler[intent] = f
}

func (s *skill) HandleSessionEndRequestFunc(f SessionEndRequestHandleFunc) {
	s.sessionEndRequestHandler = f
}

//----------------------------------------------------------------------------------------------------------------------
func (s *skill) HandleEvent(event *Event) (*ResponseBody, error) {
	var alexaResponse Response

	ctx := Context{
		memory: make(map[string]interface{}),
		Event:  event,
	}

	requestType := event.Request.RequestType()
	switch requestType {
	case LaunchRequestType:
		if s.launchRequestHandler == nil {
			return nil, errorHandlerUndefined(string(LaunchRequestType))
		}
		launchReq := event.Request.(*LaunchRequest)
		alexaResponse = s.launchRequestHandler(&ctx, launchReq)

	case IntentRequestType:
		intentReq := event.Request.(*IntentRequest)
		intent := intentReq.Intent.Name
		handler, exist := s.intentRequestHandler[intent]
		if !exist {
			return nil, errorHandlerUndefined(fmt.Sprintf("intent [%s]", intent))
		}
		alexaResponse = handler(&ctx, intentReq)

	case SessionEndedRequestType:
		if s.sessionEndRequestHandler == nil {
			return nil, errorHandlerUndefined(string(SessionEndedRequestType))
		}
		sessionEndReq := event.Request.(*SessionEndRequest)
		s.sessionEndRequestHandler(&ctx, sessionEndReq)
		return nil, nil

	default:
		if s.defaultHandler == nil {
			return nil, errorHandlerUndefined("Default")
		}
		alexaResponse = s.defaultHandler(&ctx, event)
	}

	return &ResponseBody{
		Version:           VERSION,
		SessionAttributes: ctx.memory,
		Response:          alexaResponse,
	}, nil
}

//----------------------------------------------------------------------------------------------------------------------
// handle function types
type DefaultHandleFunc = func(alexa *Context, event *Event) Response
type LaunchRequestHandleFunc = func(alexa *Context, req *LaunchRequest) Response
type IntentRequestHandleFunc = func(alexa *Context, req *IntentRequest) Response
type SessionEndRequestHandleFunc = func(alexa *Context, req *SessionEndRequest)
