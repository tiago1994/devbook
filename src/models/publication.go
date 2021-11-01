package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatedID   uint64    `json:"created_id,omitempty"`
	CreatedNick string    `json:"created_nick,omitempty"`
	Likes       uint64    `json:"likes"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (publication *Publication) Preparate() error {
	if requestError := publication.validate(); requestError != nil {
		return requestError
	}

	if requestError := publication.format(); requestError != nil {
		return requestError
	}
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("the title field is required")
	}
	if publication.Content == "" {
		return errors.New("the content field is required")
	}

	return nil
}

func (publication *Publication) format() error {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)

	return nil
}
