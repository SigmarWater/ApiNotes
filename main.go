package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"sync"

	"github.com/gin-gonic/gin"
)

type Note struct{
	ID int64 `json:"id"`
	Data time.Time `json:"data" binding:"omitempty"`
	Title string `json:"title"`
}

var mu sync.RWMutex

var notes map[int64]Note = map[int64]Note{}

func main(){
	router := gin.Default()

	router.GET("/notes", func(c *gin.Context){
		note_arr  := make([]Note, 0, len(notes))
		for _, note := range notes{
			note_arr = append(note_arr, note)
		}
		c.JSON(http.StatusOK, gin.H{"notes":note_arr})
	})

	router.GET("/notes/:id", func(c *gin.Context){
		id, err :=  strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect id"})
			return
		}

		mu.RLock()
		if note, ok := notes[id]; ok{
			c.JSON(http.StatusOK, note)
		}else{
			c.JSON(http.StatusNotFound, gin.H{"error": "not found note"})
		}
		mu.RUnlock()
	})

	router.POST("/notes", func(c *gin.Context){
		

		var note Note
		if err := c.ShouldBindJSON(&note); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad body"})
			return
		}
		
		mu.RLock()
		
		_, ok := notes[note.ID]

		mu.RUnlock()

		if ok{
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Заметка с id=%d уже есть", note.ID))
			return
		}

		if note.Title == ""{
			c.JSON(http.StatusBadRequest, gin.H{"error": "field Title is empty"})
			return 
		}
		
		mu.Lock()
		note.Data = time.Now()
		notes[note.ID] = note
		mu.Unlock()
		c.JSON(http.StatusCreated, note)
	
	})


	router.Run()
}