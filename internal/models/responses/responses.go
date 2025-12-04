package responses

import (
	"encoding/json"
)

type CV struct {
	Profile    json.RawMessage `json:"profile"`
	Projects   json.RawMessage `json:"projects"`
	Languages  json.RawMessage `json:"languages"`
	Experience json.RawMessage `json:"experience"`
	Education  json.RawMessage `json:"education"`
}
