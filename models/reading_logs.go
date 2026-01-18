package models

import "time"

type ReadingLog struct {
	LogDate     time.Time `json:"date"`
	Message     string    `json:"message"`
	CurrentPage int       `json:"currentPage"`
	Finished    bool      `json:"finished"`
}
