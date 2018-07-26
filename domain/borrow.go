package domain

import (
	"time"
)

// declare all borrow in system
type Borrow struct {
	Model
	Book_ID UUID      `json:"book_id"`
	User_ID UUID      `json:"user_id"`
	From    time.Time `sql:"default:now()" json:"from"`
	To      time.Time `json:"to"`
}
