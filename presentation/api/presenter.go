package api

import "encoding/json"

func Json(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
