package alexa

import (
	"encoding/json"
	"testing"
)

//----------------------------------------------------------------------------------------------------------------------
func TestSimpleCard_MarshalJSON(t *testing.T) {
	card, err := json.MarshalIndent(&SimpleCard{
		Title:   "SimpleCard Title",
		Content: "SimpleCard Content",
	}, "", "\t")
	if err != nil {
		t.Error(err.Error())
	}

	ans := []byte(`{
		"type": "Simple",
		"title": "SimpleCard Title",
		"content": "SimpleCard Content"
	}`)
	same, err := JSONBytesEqual(card, ans)
	if err != nil {
		t.Error(err.Error())
	}

	if !same {
		t.Error(NotSameErrMessage(string(card)))
	}
}

func TestStandardCard_MarshalJSON(t *testing.T) {
	card, err := json.MarshalIndent(&StandardCard{
		Title: "StandardCard Title",
		Text:  "StandardCard Text",
		Image: nil,
	}, "", "\t")
	if err != nil {
		t.Error(err.Error())
	}

	ans := []byte(`{
		"type": "Standard",
		"title": "StandardCard Title",
		"text": "StandardCard Text"
	}`)
	same, err := JSONBytesEqual(card, ans)
	if err != nil {
		t.Error(err.Error())
	}

	if !same {
		t.Error(NotSameErrMessage(string(card)))
	}

	// card with image
	card, err = json.MarshalIndent(&StandardCard{
		Title: "StandardCard Title",
		Text:  "StandardCard Text",
		Image: &CardImage{
			SmallImageUrl: "smallImageUrl.com",
			LargeImageUrl: "largeImageUrl.com",
		},
	}, "", "\t")
	if err != nil {
		t.Error(err.Error())
	}

	ans = []byte(`{
		"type": "Standard",
		"title": "StandardCard Title",
		"text": "StandardCard Text",
		"image": {
			"smallImageUrl": "smallImageUrl.com",
			"largeImageUrl": "largeImageUrl.com"
		}
	}`)
	same, err = JSONBytesEqual(card, ans)
	if err != nil {
		t.Error(err.Error())
	}

	if !same {
		t.Error(NotSameErrMessage(string(card)))
	}
}
