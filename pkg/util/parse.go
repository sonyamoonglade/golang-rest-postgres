package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequestBody(r *http.Request, output interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}
