package model

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	DueDate     time.Time
	Overdue     bool
	Completed   bool
}
