package dto

import "time"

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}

type URLResponse struct {
	URL        string     `json:"url"`
	Hash       string     `json:"hash"`
	CreatedAt  time.Time  `json:"created_at"`
	Expiration *time.Time `json:"expiration,omitempty"`
}
