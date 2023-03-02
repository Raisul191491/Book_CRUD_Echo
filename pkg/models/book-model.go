package models

type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	BookName    string `json:"bookname"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
