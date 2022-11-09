package models

type NewEncoding struct {
	Text string `json:"text" validate:"required"`
	Key  int    `json:"key" validate:"gte=1"`
}

type NewDecoding struct {
	Text string `json:"text" validate:"required"`
	Key  int    `json:"key" validate:"gte=1"`
}
