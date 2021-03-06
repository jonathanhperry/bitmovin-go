package models

import (
	"github.com/bitmovin/bitmovin-go/bitmovintypes"
)

type Message struct {
	Date *string                   `json:"date"`
	ID   *string                   `json:"id"`
	Type bitmovintypes.MessageType `json:"type"`
	Text *string                   `json:"text"`
}
