package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	"../entities"
)


func ShowIndexPage(c *gin.Context) {
	courses, err := entities.GetSomeCourses(0)
	if err != nil {
		log.Fatal(err)
		return
	}

	entities.Render(c,
		gin.H{"title":   "Главная страница", "payload": courses},
		"index.html")
}

func GetCourse(c *gin.Context) {
	if courseID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		course := entities.Course{}
		for i := 0; i < len(entities.FoundCourses); i++ {
			if entities.FoundCourses[i].ID == courseID {
				getCourse := entities.GetData(entities.FoundCourses[i].URLApi)
				title := getCourse.Courses[0].Title
				content := getCourse.Courses[0].Summary
				url := entities.FoundCourses[i].URL
				picture := entities.FoundCourses[i].Picture

				course = entities.Course{Title:title, Content:content, URL:url, Picture:picture,
				Host:entities.HostStepik, HostURL:entities.HostURLStepik}
			}
		}
		entities.Render(c, gin.H{"title": course.Title,
			"payload": course}, "article.html")

		/*if course, err := getCourseById(courseID); err == nil {
			render(c, gin.H{"title": course.Title,
				"payload": course}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}*/
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func SearchingRequest(c *gin.Context) {
	searchRequest := c.PostForm("research")

	coursesStepik := entities.GetStepicCourseByTitle(searchRequest)
	for i := 0; i < len(coursesStepik); i++ {
		entities.FoundCourses = append(entities.FoundCourses, coursesStepik[i])
	}

	coursesUdacity := entities.GetUdacityCourseByTitle(searchRequest)
	for i := 0; i < len(coursesUdacity); i++ {
		entities.FoundCourses = append(entities.FoundCourses, coursesUdacity[i])
	}

	entities.Render(c,
		gin.H{"title": "Результат поиска", "payload": entities.FoundCourses},
	"index.html")
}

func ShowCourseCreationPage(c *gin.Context) {
	entities.Render(c, gin.H{
		"title": "Личный кабинет"}, "personal-area.html")
}

func CreateCourse(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	id, _ := entities.GetNumberOfCourses()
	id += 1

	course := entities.Course{ Title:title, Content:content, ID:id,
	Host:entities.HostStepik, URL:entities.HostURLStepik}

	if err := course.CreateCourse(); err == nil {
		// If the article is created successfully, show success message
		entities.Render(c, gin.H{
			"title":   "Submission Successful",
			"payload": course}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
