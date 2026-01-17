package models

type List struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func (l *List) SQL() string {
	return `CREATE TABLE IF NOT EXISTS LIST(
		list_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
		
	`
}
