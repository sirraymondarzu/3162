// Filename: internal/models/questions.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// The Question model will represent a single question in our questions table
type Reservation struct {
	ReservationID int64
	UserID int64
	ReservationDate time.Time // date needs to be extracted
	Duration int8 // expected time in minutes
	PeopleCount int8 // number of people for the session
	Notes  string
	Approval bool
	CreatedAt time.Time 

    

}
// The QuestionModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type ReservationModel struct {
	DB *sql.DB
}

// The Insert() function stores a question into the questions table
func (m *ReservationModel) Insert(body string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64
	statement := 
	            `
							INSERT INTO reservations(body)
							VALUES($1)
							RETURNING question_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *ReservationModel) Get() (*Reservation, error) {
	var res Reservation

	statement := 
	            `
							SELECT question_id, body
							FROM reservations
							ORDER BY RANDOM()
							LIMIT 1
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&res.ReservationID, &res.UserID, &res.Duration, &res.PeopleCount,&res.Notes, &res.Approval,&res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}