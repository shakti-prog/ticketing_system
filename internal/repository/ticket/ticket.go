package ticket

import (
	"context"
	"ticketing_system_backend/internal/database"
)

func CreateTicket(assigneeId int64, reporterId int64, projectId int64, description string, status string, priority string) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()
	query := `INSERT INTO tickets (assignee_id,reporter_id,project_id,description,status,priority) values ($1,$2,$3,$4,$5,$6,$7)`
	_, queryError := dbpool.Exec(context.Background(), query, assigneeId, reporterId, projectId, description, status, priority)
	if queryError != nil {
		return queryError
	}
	return nil

}
