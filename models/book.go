package models

type Book struct {
	ID              uint             `json:"id"`
	BookIdentifiers []BookIdentifier `json:"bookIdentifiers"`
	Name            string           `json:"name"`
	ShortBlurb      string           `json:"shortBlurb"`
	Rating          int8             `json:"rating"`
	ProfileImage    string           `json:"profileImage"`
	TotalPageCount  int              `json:"totalPageCount"`
	Authors         []Author         `json:"authors"`
	ReadingLogs     []ReadingLog     `json:"readingLogs"`
}

type BookIdentifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Series struct {
	PartOfASeries      bool   `json:"partOfASeries"`
	Name               string `json:"name"`
	VolumeNumber       int    `json:"volumeNumber"`
	OtherBooksInSeries []Book `json:"otherBooksInSeries"`
}
