package util

import (
	"encoding/json"
	"io"
)

func ReadRequestBody(v io.ReadCloser, output interface{}) error {
	body, err := io.ReadAll(v)
	defer v.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}
