package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"../entities"
	"../auth"
)


func ShowMainPage(c *gin.Context) {
	auth.Render(c,
		gin.H{"title":   "Главная страница",},
		"index.html")
}

func GetCourse(c *gin.Context) {
	if courseID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		course := entities.Course{}
		for i := 0; i < len(entities.FoundCourses); i++ {
			if entities.FoundCourses[i].ID == courseID {
				if entities.FoundCourses[i].Host == entities.HostStepik {
					getCourse := entities.GetData(entities.FoundCourses[i].URLApi)
					title := getCourse.Courses[0].Title
					content := getCourse.Courses[0].Summary
					url := entities.FoundCourses[i].URL
					lang := entities.FoundCourses[i].Language
					skill := entities.FoundCourses[i].SkillLvl
					price := entities.FoundCourses[i].Price
					duration := getCourse.Courses[0].WorkLoad
					if duration == "" {
						duration = "-"
					}

					course = entities.Course{Title: title, Content: content, URL: url,
						Host: entities.HostStepik, HostURL: entities.HostURLStepik, Language:lang,
						SkillLvl:skill, Duration:duration, Price:price}
					break
				}
				if entities.FoundCourses[i].Host == entities.HostUdacity {

					title := entities.FoundCourses[i].Title
					content := entities.FoundCourses[i].Content
					url := entities.FoundCourses[i].URL
					lang := entities.FoundCourses[i].Language
					skill := entities.FoundCourses[i].SkillLvl
					duration := entities.FoundCourses[i].Duration
					price := entities.FoundCourses[i].Price

					course = entities.Course{Title: title, Content: content, URL: url,
						Host: entities.HostUdacity, HostURL: entities.HostURLUdacity, Language: lang,
						SkillLvl: skill, Duration: duration, Price: price}

				}

				if entities.FoundCourses[i].Host == entities.HostUdemy {
					data := entities.GetDataUdemyTest("https://www.udemy.com/api-2.0/courses/" +
						strconv.Itoa(entities.FoundCourses[i].IDUdemy) + "?fields[course]=title,headline,is_paid,num_lectures,url")

					cost := "Платно"
					if data.Is_Paid == false {
						cost = "Бесплатно"
					}
					duration := strconv.Itoa(data.Num) + " лекций"

					course = entities.Course{Title: data.Title, Content: data.Content, URL:entities.FoundCourses[i].URL,
						Host: entities.HostUdemy, HostURL: entities.HostURLUdemy,
						Language:entities.FoundCourses[i].Language,
						SkillLvl:entities.FoundCourses[i].SkillLvl, Duration: duration, Price:cost}
				}
			}
		}
		auth.Render(c, gin.H{"title": course.Title,
			"payload": course}, "course.html")

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func SearchingRequest(c *gin.Context) {
	searchRequest := c.PostForm("research")

	platforms, languages, levels, availabilities := GetParams(c)

	entities.FoundCourses = []entities.Course{}

	if len(platforms) == 0 {
		platforms = []string{ "Udacity", "Udemy", "Stepik" }
	}

	for i := 0; i < len(platforms); i++ {
		if platforms[i] == entities.HostStepik {
			if len(languages) == 2 || len(languages) == 0 {
				coursesStepik := entities.GetStepicCourseByTitle(searchRequest, "", false, 20)
				for i := 0; i < len(coursesStepik); i++ {
					entities.FoundCourses = append(entities.FoundCourses, coursesStepik[i])
				}
			} else {
				coursesStepik := entities.GetStepicCourseByTitle(searchRequest, languages[0], true, 20)
				for i := 0; i < len(coursesStepik); i++ {
					entities.FoundCourses = append(entities.FoundCourses, coursesStepik[i])
				}
			}
		}

		if platforms[i] == entities.HostUdacity {
			coursesUdacity := entities.GetUdacityCourseByTitle(searchRequest, levels, availabilities)
			for i := 0; i < len(coursesUdacity); i++ {
				entities.FoundCourses = append(entities.FoundCourses, coursesUdacity[i])
			}
		}

		if platforms[i] == entities.HostUdemy {
			coursesUdemy := entities.GetUdemyCourseByTitle(searchRequest, languages, levels, availabilities)
			for i := 0; i < len(coursesUdemy); i++ {
				entities.FoundCourses = append(entities.FoundCourses, coursesUdemy[i])
			}
		}
	}

	auth.Render(c,
		gin.H{"title": "Результат поиска", "payload": entities.FoundCourses},
	"index.html")
}

func ShowPersonalAreaPage(c *gin.Context) {
	session, _ := entities.GetLastSession()
	user,_ := entities.GetUserById(session.UserID)
	subs, _ := user.GetAllSubscriptionsByUser()

	auth.Render(c, gin.H{
		"title": "Личный кабинет", "payload": subs}, "personal-area.html")
}

func GetParams(c *gin.Context) (platforms []string, languages []string, levels []string, availabilities []string) {

	platforms_dict := []string{ "udacity", "udemy", "stepik" }
	languages_dict := []string{ "rus", "eng" }
	levels_dict := []string{ "beginner", "intermediate", "advanced" }
	availabilities_dict := []string { "free", "chargeable" }

	// Получение платформ
	for i := 0; i < len(platforms_dict); i++ {
		if c.PostForm(platforms_dict[i]) != "" {
			platforms = append(platforms, c.PostForm(platforms_dict[i]))
		}
	}
	// Получение языков
	for i := 0; i < len(languages_dict); i++ {
		if c.PostForm(languages_dict[i]) != "" {
			languages = append(languages, c.PostForm(languages_dict[i]))
		}
	}
	// Получение уровней сложности
	for i := 0; i < len(levels_dict); i++ {
		if c.PostForm(levels_dict[i]) != "" {
			levels = append(levels, c.PostForm(levels_dict[i]))
		}
	}
	// Получение способы оплаты
	for i := 0; i < len(availabilities_dict); i++ {
		if c.PostForm(availabilities_dict[i]) != "" {
			availabilities = append(availabilities, c.PostForm(availabilities_dict[i]))
		}
	}
	return
}
