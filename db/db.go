package db

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var DB *sql.DB 
var err error
var connStr string

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



