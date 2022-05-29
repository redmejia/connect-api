package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, jsonKey string, jsonVal interface{}) error {

	var data = make(map[string]interface{})
	data[jsonKey] = jsonVal

	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataByte)

	return nil
}
