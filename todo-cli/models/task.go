package models

import "time"

type Task struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}
