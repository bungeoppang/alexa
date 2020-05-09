/*
	verification codes referenced from: https://github.com/mikeflynn/go-alexa
*/
package alexa

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

var insecureSkipVerify = false

//----------------------------------------------------------------------------------------------------------------------
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-signature
func verifySignatureCertChainUrl(r *http.Request) error {
	const SignatureCertChainUrl = "SignatureCertChainUrl"

	// skip when insecure mode
	if insecureSkipVerify {
		return nil
	}

	rawUrl := r.Header.Get(SignatureCertChainUrl)
	verifyErr := errors.New(fmt.Sprintf("Invalid CertUrl: %s ", rawUrl))

	certUrl, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	// certUrl scheme is case-insensitive
	if verifyScheme(certUrl.Scheme) {
		return verifyErr
	}

	// HostName is case-insensitive
	if verifyHostName(certUrl.Hostname()) {
		return verifyErr
	}

	// if a port is specified, port is 443
	if verifyPort(certUrl.Port()) {
		return verifyErr
	}

	if verifyPathPrefix(certUrl.Path) {
		return verifyErr
	}

	//todo: cert verification

	return nil
}

func verifyScheme(scheme string) bool {
	const CertScheme = "https"
	return strings.ToLower(scheme) == CertScheme
}

func verifyHostName(hostname string) bool {
	const CertHostName = "s3.amazoneaws.com"
	return strings.ToLower(hostname) == CertHostName
}

func verifyPort(port string) bool {
	const CertPort = "443"
	return port == "" || port == CertPort
}

func verifyPathPrefix(path string) bool {
	const CertPathPrefix = "/echo.api/"
	return strings.HasPrefix(path, CertPathPrefix)
}

//----------------------------------------------------------------------------------------------------------------------
func verifyEvent(s *skill, event *Event, isDevMode bool) error {
	if !isValidApplicationId(s, event) {
		return errors.New("invalid Application id")
	}

	// skip timestamp check when devMode
	if !isDevMode {
		if !isValidTimestamp(event) {
			return errors.New("request too old to continue (>150s).")
		}
	}

	return nil
}

func isValidApplicationId(s *skill, event *Event) bool {
	return s.AppId == event.Context.System.Application.ApplicationId
}

// isValidTimestamp will parse the timestamp in the Event.Request and verify that it is in the correct
// format and is not too old. True will be returned if the timestamp is valid; false otherwise.
// https://developer.amazon.com/en-US/docs/alexa/custom-skills/host-a-custom-skill-as-a-web-service.html#check-request-timestamp
func isValidTimestamp(event *Event) bool {
	const TimeStamp = "TimeStamp"
	// todo: is there any better way to get TimeStamp from Request interface?
	reflectValue := reflect.ValueOf(event.Request)
	reflectValue = reflect.Indirect(reflectValue)
	strTimeStamp := reflectValue.FieldByName(TimeStamp).String()
	reqTimeStamp, _ := time.Parse("2019-05-13T12:34:56Z", strTimeStamp)
	return time.Since(reqTimeStamp) < time.Duration(150)*time.Second
}

//----------------------------------------------------------------------------------------------------------------------
// SetVerifyAWSCerts allows to specify whether AWS provided certs should be verified or not
func SetInsecureAWSCerts() {
	insecureSkipVerify = true
	Info("insecureSkipVerify selected, certs will not be checked")
}
