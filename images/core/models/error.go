package models

type Error struct {
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}
