package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	"fmt"
	"../entities"
)

var courses = []entities.CourseInResult{}

func ShowIndexPage(c *gin.Context) {
	courses, err := entities.GetSomeCourses(0)
	if err != nil {
		log.Fatal(err)
		return
	}

	entities.Render(c,
		gin.H{"title":   "Home Page", "payload": courses},
		"index.html")
}

func GetCourse(c *gin.Context) {
	if courseID, err := strconv.Atoi(c.Param("article_id")); err == nil {

		course := entities.Course{}
		for i := 0; i < len(courses); i++ {
			if courses[i].ID == courseID {
				apiURL := entities.GetAPICourseURLByID(courses[i].CourseID)
				getCourse := entities.GetData(apiURL)
				title := getCourse.Courses[0].Title
				content := getCourse.Courses[0].Summary
				url := courses[i].URL
				picture := courses[i].Picture

				course = entities.Course{Title:title, Content:content, URL:url, Picture:picture,
				Host:entities.Host, HostURL:entities.HostURL}
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
	fmt.Println(searchRequest)
	courses = entities.GetStepicCourseByTitle(searchRequest)

	entities.Render(c,
		gin.H{"title": "Results", "payload": courses},
	"index.html")
}

func ShowCourseCreationPage(c *gin.Context) {
	entities.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func CreateCourse(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	id, _ := entities.GetNumberOfCourses()
	id += 1

	course := entities.Course{ Title:title, Content:content, ID:id,
	Host:entities.Host, URL:entities.Url}

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
