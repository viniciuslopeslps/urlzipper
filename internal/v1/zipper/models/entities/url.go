package entities

import "time"

type URL struct {
	URL       string        `json:"url"`
	Hash      string        `json:"hash"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
}
