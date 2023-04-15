package httpmapper

import (
	"love-date/pkg/errhandling/errmsg"
	"love-date/pkg/errhandling/richerror"
	"net/http"
)

func Error(err error) (msg string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)

		msg = re.Message()

		code = mapKindToHTTPStatusCode(re.Kind())
	default:
		code = 500
	}

	if code >= 500 {
		msg = errmsg.ErrorMsgSomethingWrong
	}

	return msg, code

}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindForbidden:

		return http.StatusForbidden
	case richerror.KindNotFound:

		return http.StatusNotFound
	case richerror.KindBadRequest:

		return http.StatusBadRequest
	case richerror.KindUnexpected:

		return http.StatusInternalServerError
	default:

		return http.StatusBadRequest
	}

}
