package entities

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"strings"
)

type CourseUdacity struct {
	Courses []CourseInUdacity   `json:"courses"`
}

type CourseInUdacity struct {
	Key 		string 	`json:"key"`
	Summary 	string 	`json:"summary"`
	Title   	string  `json:"title"`
	Host    	string
	URL    		string  `json:"homepage"`
	Picture		string	`json:"image"`
}

var HostURLUdacity = "https://www.udacity.com"
var HostUdacity = "Udacity"
var myClientUdacity = &http.Client{Timeout: 100 * time.Second}

func GetUdacityCourseByTitle (title string) (courses []Course) {
	response := GetDataUdacity()
	lastID := len(FoundCourses)

	for i := 0; i < len(response.Courses); i++ {
		if IsContain(title, response.Courses[i].Title) == true {
			lastID += 1
			course := Course{ID:lastID, Title:response.Courses[i].Title,
			Content:response.Courses[i].Summary, Host:HostUdacity, HostURL:HostURLUdacity,
			URL:response.Courses[i].URL, Picture:response.Courses[i].Picture}

			courses = append(courses, course)
		}
	}
	return
}

func GetDataUdacity() (course CourseUdacity) {
	url := "https://www.udacity.com/public-api/v0/courses"
	req, err := http.NewRequest("GET", url, nil)

	resp, err := myClientUdacity.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &course)
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




