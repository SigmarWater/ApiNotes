package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SigmarWater/ApiNotes/db"

	"github.com/gin-gonic/gin"
)

func init(){
	db.ConnectDB()
}


func main() {
	router := gin.Default()

	router.GET("/notes", func(c *gin.Context) {
		notes, err := db.GetNotes()
		if err != nil{
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

		note, err := db.GetNotebyID(int(id))

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		
		c.JSON(http.StatusOK, note)
	})

	router.POST("/notes", func(c *gin.Context) {

	})

	router.Run()
}
