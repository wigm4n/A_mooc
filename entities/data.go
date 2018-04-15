package entities

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
	_ "github.com/lib/pq"
	"fmt"
	"crypto/sha1"
)

// Database instance
var db *sql.DB

// Initialize
func init() {
	file, err := ioutil.ReadFile("./db.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	type dbCredentials struct {
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		Port     string `json:"port"`
	}
	var credentials dbCredentials
	json.Unmarshal(file, &credentials)

	connection := "user=" + credentials.User +
		" password='" + credentials.Password +
		"' dbname= '" + credentials.DBName +
		"' port= '" + credentials.Port +
		"' sslmode=disable"

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
