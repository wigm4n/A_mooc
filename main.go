package main

import (
	"github.com/gin-gonic/gin"
	"./handlers"
	"./auth"
)

var router *gin.Engine

func main() {

	//gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(auth.SetUserStatus())

	// Handle the index route
	router.GET("/", handlers.ShowMainPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", auth.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", auth.EnsureNotLoggedIn(), handlers.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", auth.EnsureLoggedIn(), handlers.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", auth.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", auth.EnsureNotLoggedIn(), handlers.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handlers.GetCourse)
	}

	findRoutes := router.Group("/find")
	{
		findRoutes.POST("/searchingRequest", handlers.SearchingRequest)
	}

	personalAreaRoutes := router.Group("/personal")
	{
		personalAreaRoutes.POST("/submitting", handlers.SubmitSubscription)

		personalAreaRoutes.GET("/area", auth.EnsureLoggedIn(), handlers.ShowPersonalAreaPage)

		personalAreaRoutes.POST("/area", auth.EnsureLoggedIn(), handlers.DeleteSub)

		personalAreaRoutes.POST("/send", auth.EnsureLoggedIn(), handlers.SendEmail)
	}
}
