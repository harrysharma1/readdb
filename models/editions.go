package models

type Edition struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Number       int    `json:"number"`
	ProfileImage string `json:"profileImage"`
}
