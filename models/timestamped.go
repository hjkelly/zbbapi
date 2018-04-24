package models

import "time"

// Timestamped records creation and modification timestamps.
type Timestamped struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// SetCreationTimestamp sets creation and modification timestamps to now. This is useful when a composing model is first created.
func (t *Timestamped) SetCreationTimestamp() {
	t.Created = time.Now()
	t.Modified = time.Now()
}

// SetModificationTimestamp sets modification timestamp to now. This is useful when a composing model is updated.
func (t *Timestamped) SetModificationTimestamp() {
	t.Modified = time.Now()
}
