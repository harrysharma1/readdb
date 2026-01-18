package models

type Author struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	ProfileImage   string          `json:"profileImage"`
	ShortBio       string          `json:"shortBio"`
	AlternateNames []AlternateName `json:"alternateNames"`
}

type AlternateName struct {
	Name string `json:"name"`
}
