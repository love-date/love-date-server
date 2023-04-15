package handler

import (
	"encoding/json"
	"io"
	"love-date/pkg/errhandling/richerror"
)

func DecodeJSON(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return richerror.New("decode").WithWrapError(err).WithMessage(err.Error()).WithKind(richerror.KindBadRequest)
	}

	return nil
}
