package handlre

import (
	"encoding/json"
	"fmt"
	"io"
)

func DecodeJSON(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return fmt.Errorf("can't decoded data to JSON format %w", err)
	}

	return nil
}
