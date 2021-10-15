package models

type Post struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}
