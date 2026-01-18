package models

type Publisher struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Website      string `json:"website"`
	ProfileImage string `json:"profileImage"`
	Books        []Book `json:"books"`
}
