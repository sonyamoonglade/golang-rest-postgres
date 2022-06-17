package util

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, message map[string]interface{}) {

	w.WriteHeader(statusCode)
	mapLikeString, _ := json.Marshal(message)
	w.Write(mapLikeString)

}
