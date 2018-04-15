package entities

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type CourseUdemy struct {
	Courses []CourseInUdemy   `json:"results"`
}

type CourseInUdemy struct {
	Key     	string
	Title		string   `json:"title"`
}

var HostURLUdemy = "https://www.udemy.com"
var HostUdemy = "Udemy"
var myClientUdemy = &http.Client{Timeout: 100 * time.Second}

func GetUdemyCourseByTitle (title string) (courses []Course) {

	response := GetDataUdemy(title)

	fmt.Println(response)
	lastID := len(FoundCourses)

	for i := 0; i < len(response.Courses); i++ {
		lastID += 1
		course := Course{ID: lastID, Title: response.Courses[i].Title,
			Content: "Content", Host: HostUdemy, HostURL: HostURLUdemy,
			URL: "course URL"}

		courses = append(courses, course)
	}
	return
}

func GetDataUdemy(request string) (course CourseUdemy) {
	auth := "Basic THJyWEszZUt4WFdkaEtmNDJxeG5wUTJIZFFtYnVXT2JJZzZuWDVXUTozM2hIY0sxRkFFcnJ0ajlYZjl5aGRvM0tjM1JiTEhicWVZRDNLUkV0SlNBZVVNWG9BZlJ2a3E2TVpKY01YZDRVU0ZlU3VRTG1JbTRZWE84OXdkWk9JVkZGYWMxZHpxRjY4bW1Ob1ZLSW1vSFo2RG1xS21NRjFIdXE3bVpOc2dWdA=="

	req, err := http.NewRequest("GET", "https://www.udemy.com/api-2.0/courses/?search=" + request, nil)
	req.Header.Add("Authorization", auth)

	resp, err := myClient.Do(req)
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
