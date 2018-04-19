package models

import "time"

type Timestamped struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (t *Timestamped) SetCreationTimestamp() {
	t.Created = time.Now()
	t.Modified = time.Now()
}

func (t *Timestamped) SetModificationTimestamp() {
	t.Modified = time.Now()
}
