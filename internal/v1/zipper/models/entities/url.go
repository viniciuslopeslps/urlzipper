package entities

import (
	"encoding/json"
	"time"
)

type URL struct {
	URL       string        `json:"url"`
	Hash      string        `json:"hash"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
}

func (url URL) MarshalBinary() ([]byte, error) {
	return json.Marshal(url)
}
