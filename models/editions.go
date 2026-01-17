package models

type Edition struct {
	Name         string `json:"name"`
	Number       int    `json:"number"`
	ProfileImage string `json:"profileImage"`
}

func (e *Edition) SQL() string {
	return `CREATE TABLE IF NOT EXISTS EDITION(
		edition_id INTEGER PRIMARY KEY AUTOINCREMENT,
		book_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		number INTEGER NOT NULL,
		profile_image TEXT NOT NULL,

		FOREIGN KEY (book_id)
			REFERENCES BOOK(book_id)
			ON DELETE CASCADE
		UNIQUE(book_id, number)
	);
		
	`
}
