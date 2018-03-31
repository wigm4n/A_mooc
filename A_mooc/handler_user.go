package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"net/http"
)

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	// Obtain the POSTed username and password values
	email := c.PostForm("email")
	password := c.PostForm("password")

	isAllCorrect, _ := checkCorrectInput(email, password)

	newUser := User{}

	count, _ := getNumberOfUsers()
	count += 1

	if isAllCorrect == true {
		newUser = User{ID: count, Email: email, Password: password}
		if err := newUser.registerNewUser(); err == nil {
			// If the user is created, set the token in a cookie and log the user in
			token := generateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)

			render(c, gin.H{
				"title": "Successful registration & Login"}, "login-successful.html")

		} else {
			// If the username/password combination is invalid,
			// show the error message on the login page
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error()})

		}
	}
}

func showLoginPage(c *gin.Context) {
	render(c, gin.H{ "title": "Login"}, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("email")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}