package alexa

import (
	"errors"
	"net/http"
)

//----------------------------------------------------------------------------------------------------------------------
func errorHandlerUndefined(handlerName string) error {
	const msgPostfix = " handler is undefined"
	return errors.New(handlerName + msgPostfix)
}

//----------------------------------------------------------------------------------------------------------------------
func httpError(w http.ResponseWriter, logMsg string, errCode int) {
	if logMsg != "" {
		Info(logMsg)
	}
	http.Error(w, http.StatusText(errCode), errCode)
}
