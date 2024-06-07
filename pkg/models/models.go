package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

// defining the top level data types that our database model will use and return.
type Todos struct {
	ID      int
	Title   string
	Created time.Time
	Expires time.Time
}
