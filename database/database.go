package database

import (
	"database/sql"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int    `json:"id"`
	Fname   string `json:"fname"`
	City    string `json:"city"`
	Phone   int    `json:"phone"`
	Height  int    `json:"height`
	Married string `json:"married"`
}

type Config struct {
	DBProperties struct {
		Username      string `json:"user"`
		Password      string `json:"password"`
		Port          string `json:"port"`
		Database_name string `json:"database_name"`
		Address       string `json:"address"`
	} `yaml:"database"`
}

func GetDatabase() (*sql.DB, error) {

	db, err = sql.Open("mysql", saimohan:mohansai@tcp(localhost:3306)/saimohan
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("DB Connection Successful")
	return db, nil
}