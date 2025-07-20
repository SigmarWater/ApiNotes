package note 

import "time"

type Note struct {
	ID    int64     `json:"id"`
	Date  time.Time `json:"date" binding:"omitempty"`
	Title string    `json:"title"`
}