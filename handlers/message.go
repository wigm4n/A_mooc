package handlers

import (
	"gopkg.in/gomail.v2"
	"../entities"
	"strconv"
)

func NotifyAboutRegistration(email string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "mooc.courses.aggregator@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Уведомление о регистрации")
	m.SetBody("text/html", "Доброго времени дня, " + email + "!<p>Вы только что зарегистрировались" +
		" на <b>Агрегаторе MOOC-курсов.</b><p>Если Вы хотите получать уведомления на Ваш email, перейдите в " +
		"личный кабинет и укажите, что конкретно Вас интересует.<p>Хорошего дня,<br> Ваш Агрегатор MOOC-курсов.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "mooc.courses.aggregator@gmail.com", "12345678qwertY")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendSubEmail(email string, courses []entities.Course) {
	body := "Доброго времени дня, " + email + "!<p>Вот несколько курсов, которые могут Вам понравиться " +
		"(выборка представлена на основании Вашей подписки):<p>"
	for i := 0; i < len(courses); i++ {
		body += strconv.Itoa(i + 1) + ". " + courses[i].Host + ". " + courses[i].Title + ". URL: " + courses[i].URL + "<br>"
	}
	body += "<p>Если вы хотите отписаться от рассылки, то сделайте это в личном кабинете." +
		"<p>Хорошего дня,<br> Ваш Агрегатор MOOC-курсов."

	m := gomail.NewMessage()
	m.SetHeader("From", "mooc.courses.aggregator@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Новости по подписке")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "mooc.courses.aggregator@gmail.com", "12345678qwertY")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
