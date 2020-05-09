package alexa

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server  *httptest.Server
	alexa   *skillServer
	testUrl string
)

func TestMain(m *testing.M) {
	alexa = SkillServer("amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe")
	SetInsecureAWSCerts()
	server = httptest.NewServer(alexa)
	defer server.Close()
	testUrl = server.URL + "?_dev=true"
	os.Exit(m.Run())
}

func TestSkill_HandleLaunchRequestFunc(t *testing.T) {
	var cnt int
	alexa.HandleLaunchRequestFunc(func(alexa *Context, req *LaunchRequest) Response {
		cnt += 1
		return Tell("Hello, this is Alexa")
	})

	event := bytes.NewBufferString(launchRequestExample)
	res, err := http.Post(testUrl, "application.json", event)
	if err != nil {
		t.Error(err.Error())
	}
	if cnt != 1 {
		t.Error("Not called LaunchRequestHandler")
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	same, err := JSONBytesEqual(
		respBody, []byte(`{
		  "version": "1.0",
		  "sessionAttributes": {},
		  "response": {
			"outputSpeech": {
			  "type": "PlainText",
			  "text": "Hello, this is Alexa"
			},
			"shouldEndSession": true
		  }
	}`))
	if err != nil {
		t.Error(err.Error())
	}
	if !same {
		t.Error(NotSameErrMessage(string(respBody)))
	}
}

func TestSkill_HandleIntentRequestFunc(t *testing.T) {
	var cnt int
	alexa.HandleIntentRequestFunc("GetZodiacHoroscopeIntent", func(alexa *Context, req *IntentRequest) Response {
		cnt += 1
		return Ask("How are you?")
	})
	alexa.HandleIntentRequestFunc("mustNotAccessedIntent", func(alexa *Context, req *IntentRequest) Response {
		t.Error("invalid access")
		return nil
	})

	event := bytes.NewBufferString(intentRequestExample)
	res, err := http.Post(testUrl, "application.json", event)
	if err != nil {
		t.Error(err.Error())
	}
	if cnt != 1 {
		t.Error("Not called IntentRequestHandler")
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	same, err := JSONBytesEqual(
		respBody, []byte(`{
		  "version": "1.0",
		  "sessionAttributes": {},
		  "response": {
			"outputSpeech": {
			  "type": "PlainText",
			  "text": "How are you?"
			},
			"shouldEndSession": false
		  }
	}`))
	if err != nil {
		t.Error(err.Error())
	}
	if !same {
		t.Error(NotSameErrMessage(string(respBody)))
	}
}

func TestSkill_HandleSessionEndRequestFunc(t *testing.T) {
	var cnt int
	alexa.HandleSessionEndRequestFunc(func(alexa *Context, req *SessionEndRequest) Response {
		cnt += 1
		return Tell("Conversation Complete")
	})

	event := bytes.NewBufferString(sessionEndRequestExample)
	res, err := http.Post(testUrl, "application.json", event)
	if err != nil {
		t.Error(err.Error())
	}
	if cnt != 1 {
		t.Error("Not called SessionEndRequestHandler")
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	same, err := JSONBytesEqual(
		respBody, []byte(`{
		  "version": "1.0",
		  "sessionAttributes": {},
		  "response": {
			"outputSpeech": {
			  "type": "PlainText",
			  "text": "Conversation Complete"
			},
			"shouldEndSession": true
		  }
	}`))
	if err != nil {
		t.Error(err.Error())
	}
	if !same {
		t.Error(NotSameErrMessage(string(respBody)))
	}
}

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-types-reference.html#launchrequest-example
var launchRequestExample = `
{
  "version": "1.0",
  "session": {
    "new": true,
    "sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
    "application": {
      "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
    },
    "attributes": {},
    "user": {
      "userId": "amzn1.account.AM3B00000000000000000000000"
    }
  },
  "context": {
    "System": {
      "application": {
        "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
      },
      "user": {
        "userId": "amzn1.account.AM3B00000000000000000000000"
      },
      "device": {
        "deviceId": "string",
        "supportedInterfaces": {
          "AudioPlayer": {}
        }
      }
    },
    "AudioPlayer": {
      "offsetInMilliseconds": 0,
      "playerActivity": "IDLE"
    }
  },
  "request": {
    "type": "LaunchRequest",
    "requestId": "amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
    "timestamp": "2015-05-13T12:34:56Z",
    "locale": "ko-KR"
  }
}`

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-types-reference.html#intentrequest-example
var intentRequestExample = `
{
  "version": "1.0",
  "session": {
    "new": false,
    "sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
    "application": {
      "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
    },
    "attributes": {
      "supportedHoroscopePeriods": {
        "daily": true,
        "weekly": false,
        "monthly": false
      }
    },
    "user": {
      "userId": "amzn1.account.AM3B00000000000000000000000"
    }
  },
  "context": {
    "System": {
      "application": {
        "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
      },
      "user": {
        "userId": "amzn1.account.AM3B00000000000000000000000"
      },
      "device": {
        "supportedInterfaces": {
          "AudioPlayer": {}
        }
      }
    },
    "AudioPlayer": {
      "offsetInMilliseconds": 0,
      "playerActivity": "IDLE"
    }
  },
  "request": {
    "type": "IntentRequest",
    "requestId": " amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
    "timestamp": "2015-05-13T12:34:56Z",
    "dialogState": "COMPLETED",
    "locale": "string",
    "intent": {
      "name": "GetZodiacHoroscopeIntent",
      "confirmationStatus": "NONE",
      "slots": {
        "ZodiacSign": {
          "name": "ZodiacSign",
          "value": "virgo",
          "confirmationStatus": "NONE"
        }
      }
    }
  }
}`

// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-types-reference.html#sessionendedrequest-example
var sessionEndRequestExample = `
{
  "version": "1.0",
  "session": {
    "new": false,
    "sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
    "application": {
      "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
    },
    "attributes": {
      "supportedHoroscopePeriods": {
        "daily": true,
        "weekly": false,
        "monthly": false
      }
    },
    "user": {
      "userId": "amzn1.account.AM3B00000000000000000000000"
    }
  },
  "context": {
    "System": {
      "application": {
        "applicationId": "amzn1.echo-sdk-ams.skill.000000-d0ed-0000-ad00-000000d00ebe"
      },
      "user": {
        "userId": "amzn1.account.AM3B00000000000000000000000"
      },
      "device": {
        "supportedInterfaces": {
          "AudioPlayer": {}
        }
      }
    },
    "AudioPlayer": {
      "offsetInMilliseconds": 0,
      "playerActivity": "IDLE"
    }
  },
  "request": {
    "type": "SessionEndedRequest",
    "requestId": "amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
    "timestamp": "2015-05-13T12:34:56Z",
    "reason": "USER_INITIATED",
    "locale": "string"
  }
}`
