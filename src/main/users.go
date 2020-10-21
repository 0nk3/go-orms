package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

// User . . .
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
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
	log.Println(w, "All users End-Point Hit. . . ")
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
	log.Println(w, "New User creation End-Point Hit . . .")
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Println("Failed to open database inside AllUsers . . .")
	}
	defer db.Close()
	// lets add a user by capturing the path
	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	email := vars["email"]

	db.Create(&User{FirstName: firstname, LastName: lastname, Email: email})
	fmt.Fprintf(w, "New user added to the db . . . ")
}

// DeleteUser . . .
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Delete User End-Point Hit . . .")
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Println("Failed to open database inside AllUsers . . .")
	}
	defer db.Close()

	vars := mux.Vars(r)
	firstname := vars["firstname"]
	var user User
	db.Where("FirstName = ?", firstname).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User deleted . . . ")

}

// UpdateUser email. . .
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Update User End-Point Hit . . .")
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Println("Failed to open database inside AllUsers . . .")
	}
	defer db.Close()
	vars := mux.Vars(r)
	firstname := vars["firsname"]
	// lastname := vars["lastname"]
	email := vars["email"]

	var user User
	db.Where("firstname = ?", firstname).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "User email address updated . . . ")
}
