/*
	reference: https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
*/

package alexa

import (
	"encoding/json"
	"net/http"
)

//----------------------------------------------------------------------------------------------------------------------
func SkillServer(appId string) *skillServer {
	return &skillServer{
		skill: Skill(appId),
	}
}

//----------------------------------------------------------------------------------------------------------------------
type skillServer struct {
	*skill
}

//----------------------------------------------------------------------------------------------------------------------
// skill implements http.Handler
func (server *skillServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isDevMode := r.URL.Query().Get("_dev") != ""
	if !isDevMode {
		if err := verifySignatureCertChainUrl(r); err != nil {
			httpError(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}

	// parsing event
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		httpError(w, err.Error(), http.StatusBadRequest)
	}

	// verify event
	if err := verifyEvent(server.skill, &event, isDevMode); err != nil {
		httpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	alexaResponse, err := server.skill.HandleEvent(&event)
	if err != nil {
		httpError(w, err.Error(), http.StatusMethodNotAllowed)
	}

	b, err := json.Marshal(alexaResponse)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		Warn(err.Error())
	}
}

//----------------------------------------------------------------------------------------------------------------------
