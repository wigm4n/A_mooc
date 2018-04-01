package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"net/http"
	"../entities"
)

func GenerateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	entities.Render(c, gin.H{
		"title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	// Obtain the POSTed username and password values
	email := c.PostForm("email")
	password := c.PostForm("password")

	isAllCorrect, _ := entities.CheckCorrectInput(email, password)

	newUser := entities.User{}

	count, _ := entities.GetNumberOfUsers()
	count += 1

	if isAllCorrect == true {
		newUser = entities.User{ID: count, Email: email, Password: password}
		if err := newUser.RegisterNewUser(); err == nil {
			// If the user is created, set the token in a cookie and log the user in
			token := GenerateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)

			entities.Render(c, gin.H{
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

func ShowLoginPage(c *gin.Context) {
	entities.Render(c, gin.H{ "title": "Login"}, "login.html")
}

func PerformLogin(c *gin.Context) {
	username := c.PostForm("email")
	password := c.PostForm("password")

	if entities.IsUserValid(username, password) {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		entities.Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}