package alexa

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestQuestion(t *testing.T) {
	//rawAnswer := []byte(`{
	//	  "version": "1.0",
	//	  "sessionAttributes": {
	//		"supportedHoroscopePeriods": true
	//	  },
	//	  "response": {
	//		"outputSpeech": {
	//		  "type": "PlainText",
	//		  "text": "Today will provide you a new learning opportunity.  Stick with it and the possibilities will be endless. Can I help you with anything else?"
	//		},
	//		"card": {
	//		  "type": "Simple",
	//		  "title": "Horoscope",
	//		  "content": "Today will provide you a new learning opportunity.  Stick with it and the possibilities will be endless."
	//		},
	//		"reprompt": {
	//		  "outputSpeech": {
	//			"type": "PlainText",
	//			"text": "Can I help you with anything else?"
	//		  }
	//		},
	//		"shouldEndSession": false
	//	  }
	//}`)
	//res := Question("Today will provide you a new learning opportunity.  "+
	//	"Stick with it and the possibilities will be endless. "+
	//	"Can I help you with anything else?").
	//	Show(Card{Type: SimpleCardType,
	//		Title: "Horoscope",
	//		Content: "Today will provide you a new learning opportunity.  " +
	//			"Stick with it and the possibilities will be endless.",
	//	}).
	//	Reprompt("Can I help you with anything else?").
	//	Remember("supportedHoroscopePeriods", true)
	//
	//dres, err := json.MarshalIndent(res, "", "\t")
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//equal, err := JSONBytesEqual(dres, rawAnswer)
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//if !equal {
	//	t.Error("not equal with the answer:\n" + string(dres))
	//}

}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func NotSameErrMessage(target string) string {
	return fmt.Sprintf("not equal with the answer:\n%s", target)
}
