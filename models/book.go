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
