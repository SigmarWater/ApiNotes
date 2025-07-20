package main

import (
	"net/http"
	"strconv"
	"time"
	"github.com/SigmarWater/ApiNotes/db"
	"github.com/SigmarWater/ApiNotes/note"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/notes", func(c *gin.Context) {
		notes, err := db.GetNotes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"notes": notes})
	})

	router.GET("/notes/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect id"})
			return
		}

		note, err := db.GetNotebyID(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, note)
	})

	router.POST("/notes", func(c *gin.Context) {
		var note note.Note

		if err := c.ShouldBindJSON(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Хуево в Bind"})
			return
		}

		if note.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty title"})
			return
		}

		note.Date = time.Now()

		err := db.PostNewNote(note)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, "Note created")
		}
	})

	router.Run(":8080")
}
