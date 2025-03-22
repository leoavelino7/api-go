package entities

import (
	"github.com/google/uuid"
)

type Tweet struct {
	ID      string  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New().String(),
	}

	return &tweet
}