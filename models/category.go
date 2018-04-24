package models

import uuid "github.com/satori/go.uuid"

// Category lets you build an expected budget and categorize your actual expenses.
type Category struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name"`
	Timestamped
}
