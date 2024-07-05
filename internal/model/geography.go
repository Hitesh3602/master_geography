package model

import (
	"encoding/json"
	"time"
)

type Geography struct {
	ID        int64           `json:"id"`
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Value     string          `json:"value"`
	Metadata  json.RawMessage `json:"metadata"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
