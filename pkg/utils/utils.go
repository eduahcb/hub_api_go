package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

func ParseJSON(r *http.Request, model interface{}) error {
	return json.NewDecoder(r.Body).Decode(&model)
}

func ExpirationTime(expirationTime int) int64 {
  duration := time.Minute * time.Duration(expirationTime)
  
  expiration := time.Now().Add(duration).Unix()
  
  return expiration
}
