package handlers

import
(
	"github.com/gin-gonic/gin"
	"../entities"
	"fmt"
)

func SubmitSubscription(c *gin.Context) {
	_ = c.PostForm("research")

	result := c.PostForm("coursera")

	fmt.Print(result)

	/*coursesStepik := entities.GetStepicCourseByTitle(searchRequest)
	for i := 0; i < len(coursesStepik); i++ {
		entities.FoundCourses = append(entities.FoundCourses, coursesStepik[i])
	}

	coursesUdacity := entities.GetUdacityCourseByTitle(searchRequest)
	for i := 0; i < len(coursesUdacity); i++ {
		entities.FoundCourses = append(entities.FoundCourses, coursesUdacity[i])
	}*/

	entities.Render(c,
		gin.H{"title": "Подписка оформлена", },
		"subscription_successful.html")
}