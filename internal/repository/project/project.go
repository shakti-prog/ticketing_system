package project

import (
	"context"
	"ticketing_system_backend/internal/database"
)

func CreateProject(projectName string, createdBy int64) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()

	query := `INSERT INTO projects (project_name,created_by) values ($1, $2)`
	_, queryError := dbpool.Exec(context.Background(), query, projectName, createdBy)
	if queryError != nil {
		return queryError
	}
	return nil
}

func GetProjects() ([]string, error) {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer conn.Release()

	query := `SELECT project_name FROM projects`
	rows, queryError := dbpool.Query(context.Background(), query)
	if queryError != nil {
		return nil, queryError
	}
	defer rows.Close()

	var projectNames []string
	for rows.Next() {
		var projectName string
		if scanErr := rows.Scan(&projectName); scanErr != nil {
			return nil, scanErr
		}
		projectNames = append(projectNames, projectName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projectNames, nil
}
