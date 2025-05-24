package models

type Message struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"name"`
	Text string `json:"Text"`
}