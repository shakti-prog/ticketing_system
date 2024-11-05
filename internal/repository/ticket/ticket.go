package ticket

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"ticketing_system_backend/internal/database"
	"ticketing_system_backend/pkg/utils"
)

func CreateTicket(assigneeId int64, reporterId int64, projectId int64, description string, status string, priority string) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()
	query := `INSERT INTO tickets (assignee_id,reporter_id,project_id,description,status,priority) values ($1,$2,$3,$4,$5,$6)`
	_, queryError := dbpool.Exec(context.Background(), query, assigneeId, reporterId, projectId, description, status, priority)
	if queryError != nil {
		return queryError
	}
	return nil
}

func UpdateTicket(field string, value string, ticketNo int64) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()

	var query string
	var param interface{}

	switch field {
	case "assignee_id", "reporter_id", "project_id":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("invalid integer value")
		}
		param = int64(intValue)
	case "description", "status", "priority":
		param = value
	default:
		return errors.New("unsupported field")
	}

	query = fmt.Sprintf("UPDATE tickets SET %s = $1 WHERE ticket_no = $2", field)

	_, err := conn.Exec(context.Background(), query, param, ticketNo)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %v", err)
	}

	return nil
}

func GetSpecificTicket(ticketId int64) (map[string]interface{}, error) {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	data := make(map[string]interface{})
	if connectionErr != nil {
		return data, connectionErr
	}

	defer conn.Release()

	var ticketNo int64
	var description string
	var status string
	var priority string
	var comments []utils.CommentParams

	query := `SELECT 
 
    t.ticket_no, 
    t.description, 
    t.status, 
    t.priority, 
    ARRAY_AGG(JSONB_BUILD_OBJECT(
        'description', c.description, 
        'userId', c.user_id 
    )) AS all_comments
FROM 
    tickets AS t
LEFT JOIN 
    comments AS c ON t.id = c.ticket_id
WHERE 
    t.id = $1
GROUP BY 
    t.id, t.ticket_no, t.description, t.status, t.priority;
    `
	queryError := dbpool.QueryRow(context.Background(), query, ticketId).Scan(&ticketNo, &description, &status, &priority, &comments)
	if queryError != nil {
		return data, queryError
	}
	fmt.Println(comments);
	data["ticket_no"] = ticketNo
	data["description"] = description
	data["status"] = status
	data["priority"] = priority
	data["comments"] = comments
	return data, nil

}
