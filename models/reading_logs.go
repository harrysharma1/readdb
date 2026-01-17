package models

import "time"

type ReadingLog struct {
	LogDate     time.Time `json:"date"`
	Message     string    `json:"message"`
	CurrentPage int       `json:"currentPage"`
	Finished    bool      `json:"finished"`
}

func (r *ReadingLog) SQL() string {
	return `CREATE TABLE IF NOT EXISTS READINGLOG(
		reading_log_id INTEGER PRIMARY KEY AUTOINCREMENT,
		book_id INTEGER NOT NULL,
		log_date TEXT NOT NULL,
		message TEXT,
		current_page INTEGER CHECK(current_page>=0),
		finished INTEGER NOT NULL DEFAULT 0 CHECK(finished IN(0,1)),

		FOREIGN KEY (book_id)
			REFERENCES BOOK(book_id)
			ON DELETE CASCADE
		UNIQUE(log_date, book_id)
	);
	
	`
}
