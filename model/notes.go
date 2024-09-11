package model

type Note struct {
	ID      int    `json:"id" pg:"id"`
	Content string `json:"content" pg:"content"`
}
