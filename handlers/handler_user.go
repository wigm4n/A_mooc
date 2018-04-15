package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"net/http"
	"../entities"
	"../auth"
)

func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	auth.Render(c, gin.H{
		"title": "Регистрация"}, "register.html")
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	isAllCorrect, _ := entities.CheckCorrectInput(email, password)

	newUser := entities.User{}

	count, _ := entities.GetNumberOfUsers()
	count += 1

	if isAllCorrect == true {
		newUser = entities.User{ID: count, Email: email, Password: password}
		if err := newUser.RegisterNewUser(); err == nil {
			token := GenerateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)

			entities.DeleteAllSessions()
			newUser.CreateSession()

			NotifyAboutRegistration(email)

			auth.Render(c, gin.H{
				"title": "Успешная регистрация и авторизация"}, "login-successful.html")

		} else {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{
				"ErrorTitle":   "Ошибка регистрации",
				"ErrorMessage": err.Error()})

		}
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Ошибка регистрации",
			"ErrorMessage": "Указанный email уже зарегистрирован"})

	}
}

func ShowLoginPage(c *gin.Context) {
	auth.Render(c, gin.H{ "title": "Авторизация"}, "login.html")
}

func PerformLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if entities.IsUserValid(email, password) {
		user, _ := entities.GetUserByEmail(email)
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		entities.DeleteAllSessions()
		user.CreateSession()

		auth.Render(c, gin.H{
			"title": "Успешная авторизация"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Ошибка авторизации",
			"ErrorMessage": "Неверные данные учетной записи"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	entities.DeleteAllSessions()

	c.Redirect(http.StatusTemporaryRedirect, "/")
}