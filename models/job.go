package models

type Job struct {
	ID          string `gorm:"type:uuid;" json:"id"`
	Type        string `json:"type"`
	Url         string `json:"price"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
