package models
import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title              string   `json:"title" gorm:"not null"`
	Description        string   `json:"description" gorm:"type:text"`
	Difficulty         string   `json:"difficulty"`
	Position           string   `json:"position"`
	Type               string   `json:"type"` // New Grad, Internship, etc.
	Time               string   `json:"time"` // e.g., "Sep 2024"
	Tags               string   `json:"tags"` // Stored as a comma-separated string
	Solution           string   `json:"solution" gorm:"type:text"`
	IsProgramming      bool     `json:"isProgramming"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Company            string   `json:"company"`
	EstimatedTime      string   `json:"estimatedTime"` // e.g., "15-20 min"
}