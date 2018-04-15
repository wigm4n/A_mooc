package entities

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"strings"
	"strconv"
)

type CourseUdacity struct {
	Courses []CourseInUdacity   `json:"courses"`
	Degrees []CourseInUdacity  `json:"degrees"`
}

type CourseInUdacity struct {
	Key 			string 	`json:"key"`
	Summary 		string 	`json:"summary"`
	Title   		string  `json:"title"`
	Host    		string
	URL    			string  `json:"homepage"`
	Level 			string  `json:"level"`
	Duration   		int 	`json:"expected_duration"`
	DurationUnit	string	`json:"expected_duration_unit"`
}

var HostURLUdacity = "https://www.udacity.com"
var HostUdacity = "Udacity"
var myClientUdacity = &http.Client{Timeout: 100 * time.Second}
var languageUdacity = "English"

func GetUdacityCourseByTitle (title string, levels []string, durations []string,
	availabilities[]string) (courses []Course) {


	duration := ""
	response := GetDataUdacity()
	lastID := len(FoundCourses)

	if len(availabilities) == 0 || len(availabilities) == 2 {
		for i := 0; i < len(response.Courses); i++ {
			summary_lower := strings.ToLower(response.Courses[i].Summary)
			title_lower := strings.ToLower(title)

			if len(levels) == 0 || len(levels) == 3 {
				if strings.Contains(summary_lower, title_lower) == true {
					lastID += 1
					duration = strconv.Itoa(response.Courses[i].Duration) + " " + response.Courses[i].DurationUnit
					course := Course{ID: lastID, Title: response.Courses[i].Title,
						Content: response.Courses[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
						URL: response.Courses[i].URL, Language: languageUdacity, SkillLvl:response.Courses[i].Level,
						Price:"Бесплатно", Duration:duration}

					courses = append(courses, course)
				}
			} else {
				for j := 0; j < len(levels); j++ {
					if levels[j] == response.Courses[i].Level {
						if strings.Contains(summary_lower, title_lower) == true {
							lastID += 1
							duration = strconv.Itoa(response.Courses[i].Duration) + " " + response.Courses[i].DurationUnit
							course := Course{ID: lastID, Title: response.Courses[i].Title,
								Content: response.Courses[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
								URL: response.Courses[i].URL, Language: languageUdacity, SkillLvl: response.Courses[i].Level,
								Price: "Бесплатно", Duration:duration}

							courses = append(courses, course)
						}
					}
				}
			}
		}
		for i := 0; i < len(response.Degrees); i++ {
			summary_lower := strings.ToLower(response.Degrees[i].Summary)
			title_lower := strings.ToLower(title)

			if len(levels) == 0 || len(levels) == 3 {
				if strings.Contains(summary_lower, title_lower) == true {
					lastID += 1
					duration = strconv.Itoa(response.Degrees[i].Duration) + " " + response.Degrees[i].DurationUnit
					course := Course{ID: lastID, Title: response.Degrees[i].Title,
						Content: response.Degrees[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
						URL: response.Degrees[i].URL, Language: languageUdacity, SkillLvl:response.Degrees[i].Level,
						Price:"Платно", Duration:duration}

					courses = append(courses, course)
				}
			} else {
				for j := 0; j < len(levels); j++ {
					if levels[j] == response.Degrees[i].Level {
						if strings.Contains(summary_lower, title_lower) == true {
							lastID += 1
							duration = strconv.Itoa(response.Degrees[i].Duration) + " " + response.Degrees[i].DurationUnit
							course := Course{ID: lastID, Title: response.Degrees[i].Title,
								Content: response.Degrees[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
								URL: response.Degrees[i].URL, Language: languageUdacity, SkillLvl: response.Degrees[i].Level,
								Price: "Платно", Duration:duration}

							courses = append(courses, course)
						}
					}
				}
			}
		}
	} else if availabilities[0] == "Бесплатно" {
		for i := 0; i < len(response.Courses); i++ {
			summary_lower := strings.ToLower(response.Courses[i].Summary)
			title_lower := strings.ToLower(title)

			if len(levels) == 0 || len(levels) == 3 {
				if strings.Contains(summary_lower, title_lower) == true {
					lastID += 1
					duration = strconv.Itoa(response.Courses[i].Duration) + " " + response.Courses[i].DurationUnit
					course := Course{ID: lastID, Title: response.Courses[i].Title,
						Content: response.Courses[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
						URL: response.Courses[i].URL, Language: languageUdacity, SkillLvl:response.Courses[i].Level,
						Price:"Бесплатно", Duration:duration}

					courses = append(courses, course)
				}
			} else {
				for j := 0; j < len(levels); j++ {
					if levels[j] == response.Courses[i].Level {
						if strings.Contains(summary_lower, title_lower) == true {
							lastID += 1
							duration = strconv.Itoa(response.Courses[i].Duration) + " " + response.Courses[i].DurationUnit
							course := Course{ID: lastID, Title: response.Courses[i].Title,
								Content: response.Courses[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
								URL: response.Courses[i].URL, Language: languageUdacity, SkillLvl:response.Courses[i].Level,
								Price: "Бесплатно", Duration:duration}

							courses = append(courses, course)
						}
					}
				}
			}
		}
	} else {
		for i := 0; i < len(response.Degrees); i++ {
			summary_lower := strings.ToLower(response.Degrees[i].Summary)
			title_lower := strings.ToLower(title)

			if len(levels) == 0 || len(levels) == 3 {
				if strings.Contains(summary_lower, title_lower) == true {
					duration = strconv.Itoa(response.Degrees[i].Duration) + " " + response.Degrees[i].DurationUnit
					lastID += 1
					course := Course{ID: lastID, Title: response.Degrees[i].Title,
						Content: response.Degrees[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
						URL: response.Degrees[i].URL, Language: languageUdacity, SkillLvl:response.Degrees[i].Level,
						Price: "Платно", Duration:duration}

					courses = append(courses, course)
				}
			} else {
				for j := 0; j < len(levels); j++ {
					if levels[j] == response.Degrees[i].Level {
						if strings.Contains(summary_lower, title_lower) == true {
							lastID += 1
							duration = strconv.Itoa(response.Degrees[i].Duration) + " " + response.Degrees[i].DurationUnit
							course := Course{ID: lastID, Title: response.Degrees[i].Title,
								Content: response.Degrees[i].Summary, Host: HostUdacity, HostURL: HostURLUdacity,
								URL: response.Degrees[i].URL, Language: languageUdacity, SkillLvl: response.Degrees[i].Level,
								Price: "Платно", Duration:duration}

							courses = append(courses, course)
						}
					}
				}
			}
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
