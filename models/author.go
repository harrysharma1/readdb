package models

type Author struct {
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	ProfileImage   string          `json:"profileImage"`
	ShortBio       string          `json:"shortBio"`
	AlternateNames []AlternateName `json:"alternateNames"`
}

type AlternateName struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
