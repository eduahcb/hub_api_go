package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(r *http.Request, model interface{}) error {
	return json.NewDecoder(r.Body).Decode(&model)
}
