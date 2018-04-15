package entities

import "log"

type Session struct {
	UserID     int
}

func (user User) CreateSession() (err error) {
	statement := "INSERT INTO sessions (userID) VALUES ($1)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	return
}

func DeleteSessionByUserID(id int) (err error) {
	statement := "DELETE FROM sessions WHERE userID = $1"
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

func DeleteAllSessions() (err error) {
	statement := "DELETE FROM sessions"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return
}

func GetLastSession() (session Session, err error) {
	row, err := db.Query("SELECT * FROM sessions")
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		err = row.Scan(&session.UserID)
	}
	return
}
