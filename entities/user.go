package entities

import (
	"log"
	"errors"
	"strings"
	"fmt"
)

type User struct {
	ID int
	Email string `json:"email"`
	Password string `json:"-"`
}

func CheckCorrectInput(email string, password string) (correct bool, err error) {
	exist, _ := IsUserExist(email)
	correct = true
	if strings.TrimSpace(password) == "" {
		return false, errors.New("The password can't be empty")
	} else if exist {
		return false, errors.New("The email isn't available")
	}
	return
}

func (user *User) RegisterNewUser() (err error) {
	statement := "INSERT INTO users (ID, email, password) VALUES ($1, $2, $3)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Query(user.ID, user.Email, user.Password)
	return
}

func GetNumberOfUsers() (count int, err error) {
	count = 0
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		count += 1
	}
	return
}

func IsUserExist(email string) (exists bool, err error) {
	row, err := db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		log.Println(err)
		return
	}
	exists = row.Next()
	return
}

func GetAllUsers() (users []User, err error) {
	rows, err := db.Query("SELECT ID, email, password FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		users = append(users, user)
	}
	return
}

func IsUserValid(email string, password string) (exists bool) {
	userList, _ := GetAllUsers()
	fmt.Print(userList)
	fmt.Print(email, password)
	for _, u := range userList {
		if u.Email == email && u.Password == password {
			return true
		}
	}
	return false
}
