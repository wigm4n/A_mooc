package entities

import (
	"log"
	"fmt"
	"strings"
)

type Course struct {
	ID      	int
	Title   	string
	Content 	string
	Host		string
	HostURL		string
	URL			string
	URLApi  	string
	Price   	string
	Duration	string
	Language	string
	SkillLvl	string
}

var FoundCourses = []Course{}

func GetNumberOfCourses() (count int, err error) {
	count = 0
	rows, err := db.Query("SELECT * FROM courses")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		count += 1
	}
	return
}

func (course Course) CreateCourse() (err error) {
	statement := "INSERT INTO courses (ID, title, content, host, URL) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
		fmt.Print(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(&course.ID, &course.Title, &course.Content, &course.Host, &course.URL)
	fmt.Print(err)
	return
}

func GetCourseById(ID int) (course Course, err error) {
	err = db.QueryRow("SELECT title, content, host, url FROM courses WHERE ID = $1", ID).
		Scan(&course.Title, &course.Content, &course.Host, &course.URL)
	if err != nil {
		log.Fatal(err)
		return
	}
	course.ID = ID
	return
}


func IsContain(title string, originalTitle string) (contains bool) {
	contains = false
	arr := strings.Split(title, " ")
	arr2 := strings.Split(originalTitle, " ")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr[i] == arr2[j] {
				return true
			}
		}
	}
	return
}
