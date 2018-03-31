package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"strconv"
	"fmt"
)

func showIndexPage(c *gin.Context) {
	courses, err := getSomeCourses(0)
	if err != nil {
		log.Fatal(err)
		return
	}

	render(c,
		gin.H{"title":   "Home Page", "payload": courses},
		"index.html")
}

func getCourse(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getCourseById(articleID); err == nil {
			render(c, gin.H{"title": article.Title,
				"payload": article}, "article.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func searchingRequest(c *gin.Context) {

	searchRequest := c.PostForm("research")
	fmt.Println(searchRequest)
	courses := getStepicCourseByTitle(searchRequest)

	/*courses, err := getSomeCourses(5)
	if err != nil {
		log.Fatal(err)
		return
	}*/

	render(c,
		gin.H{"title": "Results", "payload": courses},
	"index.html")
}

func showCourseCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func createCourse(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	id, _ := getNumberOfCourses()
	id += 1

	course := Course{ Title:title, Content:content, ID:id, Host:host, URL:url}

	if err := course.createCourse(); err == nil {
		// If the article is created successfully, show success message
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": course}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
