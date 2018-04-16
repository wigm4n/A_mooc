package entities

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type CourseStepik struct {
	Courses []CourseInStepik `json:"courses"`
	Data    []CourseInResult `json:"search-results"`
	Meta    MetaStepik       `json:"meta"`
}

type MetaStepik struct {
	Page int		`json:"page"`
	Has_next bool   `json:"has_next"`
	Has_prev bool   `json:"has_previous"`
}

type CourseInStepik struct {
	ID 			int 	`json:"id"`
	Summary 	string 	`json:"summary"`
	Title   	string  `json:"title"`
	Host    	string
	URL    		string
	IsActive   	bool	`json:"is_active"`
	WorkLoad    string  `json:"workload"`
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
var HostURLStepik = "https://stepik.org/catalog"
var HostStepik = "Stepik"
var token = "jwirIzekAg1v5MSVcVbSyWe3AlWhiB"
var skillLvlStepik = "Неопределенный"
var priceStepik = "Бесплатно"

//get url terminal:
//curl -X POST -d "grant_type=client_credentials" -u"XiohmB3FE94BiQQw8huu2QzeqSF1SabDwMA9ZTvh:pjNHdemfaL01Yz0mhBgF2uVNX6YOPBepC0Jnj24E74yDTdBhBgkHSnL2ALagAeTwLaR9V4OzkdrXrHVFdwGaWTWvlQz1usDIQ81bqOxTJyqpSZJKWXOJF8yX0Z51gsvw" https://stepik.org/oauth2/token/

func GetStepicCourseByTitle(title string, lang string, flag bool, count int) (courses []Course) {
	pageNum := 0
	hasNextPage := true
	lastID := len(FoundCourses)

	for hasNextPage {
		pageNum += 1
		newURL := ""
		if flag == false {
			newURL = GetNextPageURL(pageNum, title)
		} else {
			if lang == "en" {
				newURL = GetNextPageURLEng(pageNum, title)
			} else {
				newURL = GetNextPageURLRu(pageNum, title)
			}
		}
		data := GetData(newURL)
		hasNextPage = data.Meta.Has_next

		for i := 0; i < len(data.Data); i++ {
			lastID += 1
			courseID := data.Data[i].CourseID

			courseForm := Course{ID:lastID, Title:data.Data[i].Title, Host:HostStepik,
			HostURL:HostURLStepik, URL:GetCourseURL(courseID), URLApi:GetAPICourseURLByID(courseID),
			Price:priceStepik, SkillLvl:skillLvlStepik}

			courses = append(courses, courseForm)
			if len(courses) > count {
				break
			}
		}
	}
	return
}

func GetData(url string) (course CourseStepik) {
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

func GetNextPageURLRu(number int, query string) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/api/search-results?is_popular=true&is_public=true&language=ru&page=" +
		"" + number_to_str + "&query=" + query + "&type=course"

	return
}

func GetNextPageURLEng(number int, query string) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/api/search-results?is_popular=true&is_public=true&language=en&page=" +
		"" + number_to_str + "&query=" + query + "&type=course"

	return
}

func GetNextPageURL(number int, query string) (url string) {
	number_to_str := strconv.Itoa(number)
	url = "https://stepik.org/api/search-results?is_popular=true&is_public=true&page=" +
		"" + number_to_str + "&query=" + query + "&type=course"

	return
}

func (courseInStepic CourseInStepik) StepicCourseToCourseForm(lastID int) (course Course) {
	course.ID = lastID
	course.Content = courseInStepic.Summary
	course.Title = courseInStepic.Title
	course.Host = HostStepik
	course.URL = courseInStepic.URL
	return
}
