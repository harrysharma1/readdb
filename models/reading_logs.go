package models

import "time"

type ReadingLog struct {
	ID          uint      `json:"id"`
	LogDate     time.Time `json:"date"`
	Message     string    `json:"message"`
	CurrentPage int       `json:"currentPage"`
	Finished    bool      `json:"finished"`
}
