package models

type Author struct {
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	ProfileImage   string          `json:"profileImage"`
	ShortBio       string          `json:"shortBio"`
	AlternateNames []AlternateName `json:"alternateNames"`
}

func (a *Author) SQL() string {
	return `CREATE TABLE IF NOT EXISTS AUTHOR(
		author_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		type TEXT,
		profile_image TEXT,
		short_bio TEXT	
	);

	`
}

type AlternateName struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (a *AlternateName) SQL() string {
	return `CREATE TABLE IF NOT EXISTS ALTERNATENAME(
		alternate_name_id INTEGER PRIMARY KEY AUTOINCREMENT,
		author_id INTEGER NOT NULL,
		value TEXT NOT NULL,

		FOREIGN KEY (author_id)
			REFERENCES AUTHOR(author_id)
			ON DELETE CASCADE,
		UNIQUE (author_id)
	);
	
	`
}
