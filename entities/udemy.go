package entities

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type CourseUdemy struct {
	Courses []CourseInUdemy   `json:"results"`
	Next 	string			  `json:"next"`
	Data    []string          `json:"fields"`
}

type FieldsUdemy struct {
	Title   	string 			`json:"title"`
	Content 	string			`json:"headline"`
	Is_Paid		bool			`json:"is_paid"`
	URL 		string 			`json:"url"`
	Num 		int				`json:"num_lectures"`
}

type CourseInUdemy struct {
	ID     		int		 `json:"id"`
	Title		string   `json:"title"`
	IsPaid		bool 	 `json:"is_paid"`
	URL         string   `json:"url"`
	SkillLvl	string
	Lang 		string
}

var HostURLUdemy = "https://www.udemy.com"
var HostUdemy = "Udemy"
var myClientUdemy = &http.Client{Timeout: 100 * time.Second}

func GetUdemyCourseByTitle (title string, languages []string, levels []string,
	availabilities []string) (courses []Course) {

	req_lang := ""
	req_availabilitie := ""


	if len(availabilities) == 0 || len(availabilities) == 2 {
		req_availabilitie = ""
	} else if availabilities[0] == "Бесплатно" {
		req_availabilitie = "price-free"
	} else {
		req_availabilitie = "price-paid"
	}

	if len(languages) == 0 || len(languages) == 2 {
		req_lang = ""
	} else if languages[0] == "ru" {
		req_lang = "ru"
	} else {
		req_lang = "en"
	}

	requestURL := ""
	requestURL2 := ""
	requests := []string{}
	skill := []string{}
	lang := []string{}

	if len(levels) == 0 || len(levels) == 3 {
		if req_availabilitie == "" {
			if req_lang == "" {
				requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=en&ordering=relevance"
				requestURL2 = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=ru&ordering=relevance"
			} else {
				requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=" + req_lang + "&ordering=relevance"
			}
		} else {
			if req_lang == "" {
				requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=en&ordering=relevance"
				requestURL2 = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=ru&ordering=relevance"
			} else {
				requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=" + req_lang + "&ordering=relevance"
			}
		}
		requests = append(requests, requestURL)
		skill = append(skill, "Неопределенный")
		lang = append(lang, "English")
		if requestURL2 != "" {
			requests = append(requests, requestURL2)
			skill = append(skill, "Неопределенный")
			lang = append(lang, "Русский")
		}
	} else {
		for i := 0; i < len(levels); i++ {
			req_level := ""
			if levels[i] == "advanced" {
				req_level = "expert"
			} else {
				req_level = levels[i]
			}
			if req_availabilitie == "" {
				// req_availabilitie ++
				if req_lang == "" {
					requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=en&instructional_level=" + req_level + "&ordering=relevance"
					requestURL2 = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=ru&instructional_level=" + req_level + "&ordering=relevance"
				} else {
					requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&language=" + req_lang + "&instructional_level=" + req_level + "&ordering=relevance"
				}
			} else {
				if req_lang == "" {
					requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=en&instructional_level=" + req_level + "&ordering=relevance"
					requestURL2 = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=ru&instructional_level=" + req_level + "&ordering=relevance"
				} else {
					requestURL = "https://www.udemy.com/api-2.0/courses/?page_size=100&search=" + title + "&price=" + req_availabilitie + "&language=" + req_lang + "&instructional_level=" + req_level + "&ordering=relevance"
				}
			}
			requests = append(requests, requestURL)
			skill = append(skill, req_level)
			lang = append(lang, "English")
			if requestURL2 != "" {
				requests = append(requests, requestURL2)
				skill = append(skill, req_level)
				lang = append(lang, "Русский")
			}
		}
	}


	response := CourseUdemy{}


	for i := 0; i < len(requests); i++ {
		response_temp := GetDataUdemyByURL(requests[i])

		if response_temp.Next != "null" {
			for j := 0; j < len(response_temp.Courses); j++ {
				if req_availabilitie == "price-free" {
					if response_temp.Courses[j].IsPaid == false {
						response_temp.Courses[j].SkillLvl = skill[i]
						response_temp.Courses[j].Lang = lang[i]
						response.Courses = append(response.Courses, response_temp.Courses[j])
					}
				}
				if req_availabilitie == "price-paid" {
					if response_temp.Courses[j].IsPaid == true {
						response_temp.Courses[j].SkillLvl = skill[i]
						response_temp.Courses[j].Lang = lang[i]
						response.Courses = append(response.Courses, response_temp.Courses[j])
					}
				}
				if req_availabilitie == "" {
					response_temp.Courses[j].SkillLvl = skill[i]
					response_temp.Courses[j].Lang = lang[i]
					response.Courses = append(response.Courses, response_temp.Courses[j])
				}
			}
		}
	}

	lastID := len(FoundCourses)

	for i := 0; i < len(response.Courses); i++ {
		lastID += 1
		price := "Бесплатно"
		if response.Courses[i].IsPaid == true {
			price = "Платно"
		}
		if response.Courses[i].SkillLvl == "expert" {
			response.Courses[i].SkillLvl = "advanced"
		}
		course := Course{ID: lastID, Title: response.Courses[i].Title, Host: HostUdemy, HostURL: HostURLUdemy,
			URL: "https://www.udemy.com" + response.Courses[i].URL, Price:price, SkillLvl:response.Courses[i].SkillLvl,
			IDUdemy:response.Courses[i].ID, Language:response.Courses[i].Lang}

		courses = append(courses, course)
	}
	return
}

func GetDataUdemyByURL(request string) (course CourseUdemy) {
	auth := "Basic THJyWEszZUt4WFdkaEtmNDJxeG5wUTJIZFFtYnVXT2JJZzZuWDVXUTozM2hIY0sxRkFFcnJ0ajlYZjl5aGRvM0tjM1JiTEhicWVZRDNLUkV0SlNBZVVNWG9BZlJ2a3E2TVpKY01YZDRVU0ZlU3VRTG1JbTRZWE84OXdkWk9JVkZGYWMxZHpxRjY4bW1Ob1ZLSW1vSFo2RG1xS21NRjFIdXE3bVpOc2dWdA=="
	req, err := http.NewRequest("GET", request, nil)
	req.Header.Add("Authorization", auth)

	resp, err := myClientUdemy.Do(req)
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

func GetDataUdemyTest(request string) (course FieldsUdemy) {
	auth := "Basic THJyWEszZUt4WFdkaEtmNDJxeG5wUTJIZFFtYnVXT2JJZzZuWDVXUTozM2hIY0sxRkFFcnJ0ajlYZjl5aGRvM0tjM1JiTEhicWVZRDNLUkV0SlNBZVVNWG9BZlJ2a3E2TVpKY01YZDRVU0ZlU3VRTG1JbTRZWE84OXdkWk9JVkZGYWMxZHpxRjY4bW1Ob1ZLSW1vSFo2RG1xS21NRjFIdXE3bVpOc2dWdA=="
	req, err := http.NewRequest("GET", request, nil)
	req.Header.Add("Authorization", auth)

	resp, err := myClientUdemy.Do(req)
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
