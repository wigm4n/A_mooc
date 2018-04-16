package handlers

import
(
	"github.com/gin-gonic/gin"
	"../entities"
	"../auth"
	"time"
	"strconv"
	"log"
	"strings"
	"fmt"
)

func SubmitSubscription(c *gin.Context) {
	platforms := ""
	languages := ""
	levels := ""
	availabilities := ""

	platforms_dict := []string{ "udacity", "udemy", "stepik"}
	languages_dict := []string{ "rus", "eng" }
	levels_dict := []string{ "beginner", "intermediate", "advanced" }
	availabilities_dict := []string { "free", "chargeable" }

	isEmpty := true
	for i := 0; i < len(platforms_dict); i++ {
		if c.PostForm(platforms_dict[i]) != "" {
			if len(platforms) != 0 {
				platforms += ", " + c.PostForm(platforms_dict[i])
			} else {
				platforms += c.PostForm(platforms_dict[i])
				isEmpty = false
			}
		}
	}
	if isEmpty {
		platforms = "Любая"
	} else {
		isEmpty = true
	}


	for i := 0; i < len(languages_dict); i++ {
		if c.PostForm(languages_dict[i]) != "" {
			if len(languages) != 0 {
				languages += ", " + c.PostForm(languages_dict[i])
			} else {
				languages += c.PostForm(languages_dict[i])
				isEmpty = false
			}
		}
	}
	if isEmpty {
		languages = "Любой"
	} else {
		isEmpty = true
	}

	for i := 0; i < len(levels_dict); i++ {
		if c.PostForm(levels_dict[i]) != "" {
			if len(levels) != 0 {
				levels += ", " + c.PostForm(levels_dict[i])
			} else {
				levels += c.PostForm(levels_dict[i])
				isEmpty = false
			}
		}
	}
	if isEmpty {
		levels = "Любой"
	} else {
		isEmpty = true
	}

	for i := 0; i < len(availabilities_dict); i++ {
		if c.PostForm(availabilities_dict[i]) != "" {
			if len(availabilities) != 0 {
				availabilities += ", " + c.PostForm(availabilities_dict[i])
			} else {
				availabilities += c.PostForm(availabilities_dict[i])
				isEmpty = false
			}
		}
	}
	if isEmpty {
		availabilities = "Любая"
	} else {
		isEmpty = true
	}

	teg := c.PostForm("teg")
	frequency, _ := strconv.Atoi(c.PostForm("selectFrequency"))
	date := time.Now()

	session, _ := entities.GetLastSession()
	user, _ := entities.GetUserById(session.UserID)
	id, _ := entities.GetNumberOfSubscriptions(session.UserID)

	newSubscription := entities.Subscription{ ID:id, UserID:session.UserID, Teg:teg, Platforms:platforms,
	Languages:languages, Levels:levels, Availabilities:availabilities,
	Date:date, Frequency:frequency}

	err := user.CreateSubscription(newSubscription)

	fmt.Println(err)

	auth.Render(c,
		gin.H{"title": "Подписка оформлена", },
		"subscription-successful.html")
}

func DeleteSub(c *gin.Context) {
	sub_id, _ := strconv.Atoi(c.PostForm("deleteId"))
	err := entities.DeleteSubscriptionById(sub_id)
	if err != nil {
		log.Println(err)
	}

	auth.Render(c,
		gin.H{"title": "Подписка удалена", },
		"sub_deleting_successful.html")
}

func SendEmail(c *gin.Context) {

	sub_id, _ := strconv.Atoi(c.PostForm("sendId"))
	sub, err := entities.GetSubscriptionsByID(sub_id)
	if err != nil {
		log.Println(err)
	}

	user, _ := entities.GetUserById(sub.UserID)
	email := user.Email
	string_request := sub.Teg
	platforms := []string{}
	languages := []string{}
	levels := []string{}
	costs := []string{}

	if sub.Platforms == "Любая" {
		platforms = []string{ "Udacity", "Udemy", "Stepik"}
	} else {
		arr := strings.Split(sub.Platforms, ", ")
		for i := 0; i < len(arr); i++ {
			platforms = append(platforms, arr[i])
		}
	}

	if sub.Languages == "Любой" {
		languages = []string{ "en", "ru"}
	} else {
		if sub.Languages == "Английский" {
			languages = append(languages, "en")
		} else {
			languages = append(languages, "ru")
		}
	}

	if sub.Levels == "Любой" {
		levels = []string{ "beginner", "intermediate", "advanced"}
	} else {
		arr := strings.Split(sub.Levels, ", ")
		for i := 0; i < len(arr); i++ {
			if arr[i] == "Начинающий" {
				levels = append(levels, "beginner")
			}
			if arr[i] == "Промежуточный" {
				levels = append(levels, "intermediate")
			}
			if arr[i] == "Продвинутый" {
				levels = append(levels, "advanced")
			}
		}
	}

	if sub.Availabilities == "Любая" {
		costs = []string{ "Платно", "Бесплатно"}
	} else {
		arr := strings.Split(sub.Availabilities, ", ")
		for i := 0; i < len(arr); i++ {
			costs = append(costs, arr[i])
		}
	}

	var coursesForSub = []entities.Course{}

	for k := 0; k < len(platforms); k++ {
		// Stepik
		if platforms[k] == "Stepik" {
			if len(languages) == 2 {
				coursesStepik := entities.GetStepicCourseByTitle(string_request, "", false, 3)
				if len(coursesStepik) < 3 {
					for i := 0; i < len(coursesStepik); i++ {
						coursesForSub = append(coursesForSub, coursesStepik[i])
					}
				} else {
					for i := 0; i < 3; i++ {
						coursesForSub = append(coursesForSub, coursesStepik[i])
					}
				}
			} else {
				coursesStepik := entities.GetStepicCourseByTitle(string_request, languages[0], true, 3)
				if len(coursesStepik) < 3 {
					for i := 0; i < len(coursesStepik); i++ {
						coursesForSub = append(coursesForSub, coursesStepik[i])
					}
				} else {
					for i := 0; i < 3; i++ {
						coursesForSub = append(coursesForSub, coursesStepik[i])
					}
				}
			}
		}

		// Udacity
		if platforms[k] == "Udacity" {
			coursesUdacity := entities.GetUdacityCourseByTitle(string_request, levels, costs)
			if len(coursesUdacity) < 3 {
				for i := 0; i < len(coursesUdacity); i++ {
					coursesForSub = append(coursesForSub, coursesUdacity[i])
				}
			} else {
				for i := 0; i < 3; i++ {
					coursesForSub = append(coursesForSub, coursesUdacity[i])
				}
			}
		}

		// Udemy
		if platforms[k] == "Udemy" {
			coursesUdemy := entities.GetUdemyCourseByTitle(string_request, languages, levels, costs)
			if len(coursesUdemy) < 3 {
				for i := 0; i < len(coursesUdemy); i++ {
					coursesForSub = append(coursesForSub, coursesUdemy[i])
				}
			} else {
				for i := 0; i < 3; i++ {
					coursesForSub = append(coursesForSub, coursesUdemy[i])
				}
			}
		}
	}

	SendSubEmail(email, coursesForSub)

	auth.Render(c,
		gin.H{"title": "Сообщение отправлено", },
		"sub_email_successful.html")
}