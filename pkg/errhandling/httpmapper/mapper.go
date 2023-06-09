package httpmapper

import (
	"fmt"
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

		fmt.Println("err message", msg)
		fmt.Println("err kind", code)
		fmt.Println()

		if code >= 500 {
			msg = errmsg.ErrorMsgSomethingWrong
		}

		return msg, code
	default:

		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindForbidden:

		return http.StatusForbidden
	case richerror.KindNotFound:

		return http.StatusNotFound
	case richerror.KindBadRequest:

		return http.StatusBadRequest

	case richerror.KindUnauthorized:

		return http.StatusUnauthorized
	case richerror.KindUnexpected:

		return http.StatusInternalServerError

	case richerror.KindNoContent:

		return http.StatusNoContent
	default:

		return http.StatusBadRequest
	}

}
