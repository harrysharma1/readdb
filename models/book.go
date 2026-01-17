package models

type Book struct {
	BookIDs        []BookID     `json:"bookIds"`
	Name           string       `json:"name"`
	Editions       []Edition    `json:"editions"`
	ShortBlurb     string       `json:"shortBlurb"`
	Rating         int8         `json:"rating"`
	ProfileImage   string       `json:"profileImage"`
	TotalPageCount int          `json:"totalPageCount"`
	Authors        []Author     `json:"authors"`
	ReadingLogs    []ReadingLog `json:"readingLogs"`
}

func (b *Book) SQL() string {
	return `CREATE TABLE IF NOT EXISTS BOOK(
		book_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		short_blurb TEXT NOT NULL,
		rating INTEGER CHECK (rating BETWEEN 1 AND 5),
		profile_image TEXT NOT NULL,
		total_page_count INTEGER
	);

	CREATE TABLE IF NOT EXISTS BOOKAUTHOR(
		book_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,

		PRIMARY KEY(book_id, author_id)

		FOREIGN KEY (author_id)
			REFERENCES AUTHOR(author_id)
		
		FOREIGN KEY (book_id)
			REFERENCES BOOK(book_id)

	);

	CREATE TABLE IF NOT EXISTS BOOKLISTS(
		book_id INTEGER NOT NULL,
		list_id INTEGER NOT NULL,

		PRIMARY KEY (book_id, list_id)

		FOREIGN KEY(list_id)
			REFERENCES LIST(list_id)

		FOREIGN KEY(book_id)
			REFERENCES BOOK(book_id)
	);

	`
}

type BookID struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (b *BookID) SQL() string {
	return `CREATE TABLE IF NOT EXISTS BOOKIDENTIFIERS(
		book_identifiers_id INTEGER PRIMARY KEY AUTOINCREMENT,
		book_id INTEGER NOT NULL,
		type TEXT NOT NULL,
		value TEXT NOT NULL,

		FOREIGN KEY (book_id)
			REFERENCES BOOK(book_id)
			ON DELETE CASCADE
		UNIQUE(book_id, value)
	);
	
	`
}
