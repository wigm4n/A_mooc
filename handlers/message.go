package handlers

import (
	"gopkg.in/gomail.v2"
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
