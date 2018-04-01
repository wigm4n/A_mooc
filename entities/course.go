package entities

import (
	"log"
	"fmt"
)

type Course struct {
	ID      int
	Title   string
	Content string
	Host	string
	HostURL	string
	URL		string
	Picture	string
}

func GetSomeCourses(len int) (courses[]Course, err error) {
	rows, err := db.Query("SELECT ID, title, content, host, url FROM courses")
	if err != nil {
		log.Fatal(err)
		return
	}
	count := 0
	for rows.Next() {
		if count == len {
			break
		}
		var course Course
		err = rows.Scan(&course.ID, &course.Title, &course.Content, &course.Host, &course.URL)
		courses = append(courses, course)
		count += 1
	}
	return
}

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
