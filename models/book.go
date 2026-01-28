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
	BookBuyLinks    []BookBuyLink    `json:"bookBuyLinks"`
}

type BookIdentifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type BookBuyLink struct {
	Retailer string `json:"retailer"`
	URL      string `json:"url"`
}

type Series struct {
	ID                 uint   `json:"id"`
	PartOfASeries      bool   `json:"partOfASeries"`
	Name               string `json:"name"`
	OtherBooksInSeries []Book `json:"otherBooksInSeries"`
}
