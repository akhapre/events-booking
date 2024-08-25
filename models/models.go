package models

import (
	"events-booking/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, ?)`
	preparedStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer preparedStatement.Close()
	result, err := preparedStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	lastInsertedId, _ := result.LastInsertId()
	e.ID = lastInsertedId
	return nil
}

func GetAllEvents() (*[]Event, error) {
	query := `SELECT * from events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return &events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * from events where id=?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event *Event) UpdateEvent() error {
	query := `UPDATE events 
		SET name=?, 
		description=?, 
		location=?, 
		dateTime=?
		where id=?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}
