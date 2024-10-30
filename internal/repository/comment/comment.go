package Comment

import (
	"context"
	"ticketing_system_backend/internal/database"
)

func CreateComment(ticketId int64, userId int64, description string) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()

	query := `INSERT INTO comments (ticket_id,user_id,description) values ($1, $2, $3)`
	_, queryError := dbpool.Exec(context.Background(), query, ticketId, userId, description)
	if queryError != nil {
		return queryError
	}
	return nil
}
