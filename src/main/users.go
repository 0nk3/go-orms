package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// User . . .
type User struct {
	gorm.Model
	firstname string
	lastname  string
	email     string
}

// InitialMigration  . . .
func InitialMigration() {
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Println(err.Error())
		panic("Failed to open database")
	}
	log.Println("Connected to the database . . . ")
	defer db.Close()
	db.AutoMigrate(&User{})
}

// AllUsers . . .
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users End-Point Hit. . . ")
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Println("Failed to open database inside AllUsers . . .")
	}
	defer db.Close()
	// lets define a temp storage location for ou user

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

// NewUser . . .
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User creation End-Point Hit . . .")
}

// DeleteUser . . .
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User End-Point Hit . . .")
}

// UpdateUser . . .
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User End-Point Hit . . .")
}
