package entities

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type CourseStepic struct {
	Courses []CourseInStepic   `json:"courses"`
	Data []CourseInResult 	   `json:"search-results"`
	Meta    MetaStepic         `json:"meta"`
}

type MetaStepic struct {
	Page int		`json:"page"`
	Has_next bool   `json:"has_next"`
	Has_prev bool   `json:"has_previous"`
}

type CourseInStepic struct {
	ID 			int 	`json:"id"`
	Summary 	string 	`json:"summary"`
	Title   	string  `json:"title"`
	Host    	string
	URL    		string
	IsActive   	bool	`json:"is_active"`
}

type CourseInResult struct {
	ID 			int
	CourseID	int 	`json:"course"`
	Title		string	`json:"course_title"`
	Host		string
	URL         string
	URLapi		string
	Picture		string  `json:"course_cover"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}
var Url = "https://stepik.org/api/courses/67"
var HostURL = "https://stepik.org/catalog"
var Host = "Stepik"
var token = "jwirIzekAg1v5MSVcVbSyWe3AlWhiB"

func stepikWork(req string) {
	//get url terminal:
	//curl -X POST -d "grant_type=client_credentials" -u"XiohmB3FE94BiQQw8huu2QzeqSF1SabDwMA9ZTvh:pjNHdemfaL01Yz0mhBgF2uVNX6YOPBepC0Jnj24E74yDTdBhBgkHSnL2ALagAeTwLaR9V4OzkdrXrHVFdwGaWTWvlQz1usDIQ81bqOxTJyqpSZJKWXOJF8yX0Z51gsvw" https://stepik.org/oauth2/token/

	/*courses := GetStepicCourseByTitle(req)
	for i:= 0; i < len(courses); i++ {
		courses[i].CreateCourse()
	}*/
}

func GetStepicCourseByTitle(title string) (courses []CourseInResult) {
	pageNum := 0
	hasNextPage := true
	lastID, _ := GetNumberOfCourses()

	for hasNextPage {
		pageNum += 1
		newURL := GetNextPageURL(pageNum, title)
		data := GetData(newURL)
		hasNextPage = data.Meta.Has_next

		for i := 0; i < len(data.Data); i++ {

			//courseTitle := data.Data[i].Title

			//if IsContain(title, courseTitle) == true
			courseTitle := data.Data[i].Title
			courseID := data.Data[i].CourseID
			URLapi := GetAPICourseURLByID(courseID)
			URL := GetCourseURL(courseID)
			picture := data.Data[i].Picture

			lastID += 1

			courseForm := CourseInResult{ID:lastID, CourseID:courseID,
			Title:courseTitle, Host:Host, URL:URL, URLapi:URLapi, Picture:picture}

			courses = append(courses, courseForm)

			fmt.Println("Added:", courseForm)
		}
	}
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

func GetData(url string) (course CourseStepic) {
	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", bearer)

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

func GetCourseURL(number int) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/course/" + number_to_str
	return
}

func GetAPICourseURLByID(number int) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/api/courses/" + number_to_str
	return
}

func GetNextPageURL(number int, query string) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/api/search-results?is_popular=true&is_public=true&page=" +
		"" + number_to_str + "&query=" + query + "&type=course"

	return
}

func (courseInStepic CourseInStepic) StepicCourseToCourseForm(lastID int) (course Course, err error) {
	course.ID = lastID
	course.Content = courseInStepic.Summary
	course.Title = courseInStepic.Title
	course.Host = Host
	course.URL = courseInStepic.URL
	return
}
