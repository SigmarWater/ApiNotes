package db

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/SigmarWater/ApiNotes/note"
)


var DB *sql.DB 
var err error
var connStr string
var LastNoteId *int64 = new(int64) 

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=notes sslmode=disable", host, port, user, password)
}




func ConnectDB(){

	DB, err = sql.Open("postgres", connStr)

	if err != nil{
		log.Println("error open DB")
		return 
	}

	err = DB.Ping()
	if err != nil{
		log.Println(err.Error())
		return
	}

	log.Println("succses connection")
}

func GetNotes() ([]note.Note, error){
	notes := make([]note.Note, 0, 10)

	query := "SELECT title, date FROM note"
	rows, err := DB.Query(query)
	if err != nil{
		log.Println("error Query")
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var note note.Note
		err = rows.Scan(&note.Title, &note.Date)
		if err != nil{
			log.Println("error binding")
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {  // (5)
    	log.Println(err)
		return nil, err
	}

	return notes, nil 
}

// func GetNotebyID(id int64) (note.Note, error){
// 	var note note.Note

// 	query := "SELECT title, date from note where id=$1"

// 	row  := DB.QueryRow(query, id)

// 	err = row.Scan(&note.Date, &note.Title)

// 	if err == sql.ErrNoRows{
// 		log.Printf("no note with id = %d", id)
// 		return note, err
// 	}else if err != nil{
// 		log.Println(err.Error())
// 		return note, err
// 	}else{
// 		return note, err
// 	}
// }

	

func PostNewNote(note note.Note) error{

	query := "INSERT INTO note(date, title) VALUES ($1, $2)"

	_, err := DB.Exec(query, note.Date, note.Title)
	if err != nil{
		return err
	}
	
	return nil
}
