package models

type Message struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"name"`
	Text string `json:"Text"`
}