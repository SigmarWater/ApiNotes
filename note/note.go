package note 

import "time"

type Note struct {
	Date  time.Time `json:"date" binding:"omitempty"`
	Title string    `json:"title"`
}