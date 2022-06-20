package util

import (
	"encoding/json"
	"io"
)

func ReadRequestBody(rc io.ReadCloser, output interface{}) error {
	body, err := io.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}
