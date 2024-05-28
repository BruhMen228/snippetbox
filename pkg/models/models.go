package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: не найдено нужной записи")

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}