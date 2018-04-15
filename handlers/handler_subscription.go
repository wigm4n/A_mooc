package handlers

import
(
	"github.com/gin-gonic/gin"
	"../entities"
	"../auth"
	"time"
	"strconv"
	"fmt"
)

func SubmitSubscription(c *gin.Context) {
	platforms := ""
	languages := ""
	levels := ""
	durations := ""
	availabilities := ""

	platforms_dict := []string{ "coursera", "udacity", "edx", "udemy", "stepik"}
	languages_dict := []string{ "russian", "english" }
	levels_dict := []string{ "beginner", "intermediate", "advanced" }
	durations_dict := []string { "less_m", "one_three_m", "more_m" }
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

	for i := 0; i < len(durations_dict); i++ {
		if c.PostForm(durations_dict[i]) != "" {
			if len(durations) != 0 {
				durations += ", " + c.PostForm(durations_dict[i])
			} else {
				durations += c.PostForm(durations_dict[i])
				isEmpty = false
			}
		}
	}
	if isEmpty {
		durations = "Любая"
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
	s := c.PostForm("selectFrequency")
	fmt.Println(s)
	date := time.Now()

	session, _ := entities.GetLastSession()
	user, _ := entities.GetUserById(session.UserID)
	id, _ := entities.GetNumberOfSubscriptions(session.UserID)

	newSubscription := entities.Subscription{ ID:id, UserID:session.UserID, Teg:teg, Platforms:platforms,
	Languages:languages, Levels:levels, Durations:durations, Availabilities:availabilities,
	Date:date, Frequency:frequency}

	user.CreateSubscription(newSubscription)

	auth.Render(c,
		gin.H{"title": "Подписка оформлена", },
		"subscription-successful.html")
}