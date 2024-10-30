package signup

import (
	"context"
	"ticketing_system_backend/internal/database"
	"ticketing_system_backend/pkg/utils"
)

func SignUp(userName string, userEmail string, userPassword string) error {
	dbpool := database.ConnectDB()
	conn, connectionErr := dbpool.Acquire(context.Background())
	if connectionErr != nil {
		return connectionErr
	}
	defer conn.Release()
	hashedPassword, hashingError := utils.GenerateHash(userPassword)
	if hashingError != nil {
		return hashingError
	}
	query := `INSERT INTO users (name,email,password) values ($1, $2, $3)`
	_, queryError := dbpool.Exec(context.Background(), query, userName, userEmail, hashedPassword)
	if queryError != nil {
		return queryError
	}
	return nil
}
