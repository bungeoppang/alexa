package alexa

import "encoding/json"

//----------------------------------------------------------------------------------------------------------------------
type CardType string

const (
	SimpleCardType                   CardType = "Simple"
	StandardCardType                 CardType = "Standard"
	LinkAccountCardType              CardType = "LinkAccount"
	AskForPermissionsConsentCardType CardType = "AskForPermissionsConsent"
)

//----------------------------------------------------------------------------------------------------------------------
// Card type
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/request-and-response-json-reference.html#outputspeech-object
type Card interface {
	AlexaCard()
}

//----------------------------------------------------------------------------------------------------------------------
// type SimpleCard struct
type SimpleCard struct {
	Title   string
	Content string
}

func (card *SimpleCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type    CardType `json:"type"`
		Title   string   `json:"title,omitempty"`
		Content string   `json:"content,omitempty"`
	}{
		Type:    SimpleCardType,
		Title:   card.Title,
		Content: card.Content,
	})
}

//----------------------------------------------------------------------------------------------------------------------
// type StandardCard struct
type StandardCard struct {
	Title string
	Text  string
	Image *CardImage
}

func (card *StandardCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  CardType   `json:"type"`
		Title string     `json:"title,omitempty"`
		Text  string     `json:"text,omitempty"`
		Image *CardImage `json:"image,omitempty"`
	}{
		Type:  StandardCardType,
		Title: card.Title,
		Text:  card.Text,
		Image: card.Image,
	})
}

type CardImage struct {
	SmallImageUrl string `json:"smallImageUrl"`
	LargeImageUrl string `json:"largeImageUrl"`
}

//----------------------------------------------------------------------------------------------------------------------
// type LinkAccountCard struct
// todo: specify LinkAccountCard
type LinkAccountCard struct {
}

func (card *LinkAccountCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type CardType `json:"type"`
	}{
		Type: LinkAccountCardType,
	})
}

//----------------------------------------------------------------------------------------------------------------------
// type AskForPermissionsConsentCard struct
// todo: specify AskForPermissionsConsentCard
type AskForPermissionsConsentCard struct {
}

func (card *AskForPermissionsConsentCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type CardType `json:"type"`
	}{
		Type: AskForPermissionsConsentCardType,
	})
}

//----------------------------------------------------------------------------------------------------------------------
// implementation of alexa.Card interface
func (SimpleCard) AlexaCard()                   {}
func (StandardCard) AlexaCard()                 {}
func (LinkAccountCard) AlexaCard()              {}
func (AskForPermissionsConsentCard) AlexaCard() {}
