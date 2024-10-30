package login

import (
	"context"
	"database/sql"
	"fmt"
	"ticketing_system_backend/internal/database"
	"ticketing_system_backend/pkg/utils"
)

func Login(userEmail string, userPassword string) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()
	query := `
		SELECT name,password FROM users WHERE email = $1;
	`
	var name string
	var password string

	queryError := dbpool.QueryRow(context.Background(), query, userEmail).Scan(&name, &password)
	if queryError != nil {
		if queryError == sql.ErrNoRows {
			return fmt.Errorf("no user found with this email")
		}
		return fmt.Errorf("failed to retrieve user: %w", queryError)
	}
	if !utils.VerifyHash(userPassword, password) {
		return fmt.Errorf("wrong password")
	}
	return nil

}
