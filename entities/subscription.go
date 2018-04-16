package entities

import (
	"time"
	"fmt"
)

type Subscription struct {
	ID				int
	UserID 			int
	Teg 			string
	Platforms		string
	Languages		string
	Levels			string
	Availabilities	string
	Date			time.Time
	Frequency		int
}

func (user User) CreateSubscription(subscription Subscription) (err error) {
	statement := "INSERT INTO subscriptions (id, userID, teg, platforms, languages, levels," +
		" availabilities, date, frequency) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(subscription.ID, user.ID, subscription.Teg, subscription.Platforms,
		subscription.Languages, subscription.Levels, subscription.Availabilities,
			subscription.Date, subscription.Frequency)
	return
}

func (user User) GetAllSubscriptionsByUser() (subscriptions []Subscription, err error) {
	rows, err := db.Query("SELECT id, userID, teg, platforms, languages, levels," +
		" availabilities, date, frequency FROM subscriptions WHERE userID = $1", user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var subscription Subscription
		err = rows.Scan(&subscription.ID, &subscription.UserID, &subscription.Teg, &subscription.Platforms,
			&subscription.Languages, &subscription.Levels, &subscription.Availabilities,
				&subscription.Date, &subscription.Frequency)
		subscriptions = append(subscriptions, subscription)
	}
	return
}

func GetSubscriptionsByID(id int) (subscription Subscription, err error) {
	rows, err := db.Query("SELECT id, userID, teg, platforms, languages, levels,"+
		" availabilities, date, frequency FROM subscriptions WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&subscription.ID, &subscription.UserID, &subscription.Teg, &subscription.Platforms,
			&subscription.Languages, &subscription.Levels, &subscription.Availabilities,
			&subscription.Date, &subscription.Frequency)
	}
	return
}

func DeleteSubscriptionById(id int) (err error) {
	statement := "DELETE FROM subscriptions WHERE id = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return
}

func (user User) DeleteAllSubscriptions() (err error) {
	statement := "DELETE FROM subscriptions WHERE userID = $1"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		return err
	}
	return
}

func GetNumberOfSubscriptions(id int) (count int, err error) {
	count = 0
	rows, err := db.Query("SELECT * FROM subscriptions")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		count += 1
	}
	return
}
