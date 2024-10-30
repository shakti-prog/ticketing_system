package migrations

import (
	"context"
	"fmt"
	"log"
	"ticketing_system_backend/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTables() {
	dbpool := database.ConnectDB()

	createTables(dbpool)

}

func createTables(dbpool *pgxpool.Pool) {
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Error in establishing connection")
	}
	defer conn.Release()
	queries := []string{
		// Create users table
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		// Create projects table
		`CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			project_name VARCHAR(100) NOT NULL,
			created_by INTEGER ,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE SET NULL
		);`,

		// Create tickets table
		`CREATE TABLE IF NOT EXISTS tickets (
			id SERIAL PRIMARY KEY,
			ticket_no INTEGER UNIQUE NOT NULL,
			assignee_id INTEGER,
			reporter_id INTEGER,
			project_id INTEGER NOT NULL,
			description TEXT,
			status VARCHAR(50),
			priority VARCHAR(50),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_assignee FOREIGN KEY (assignee_id) REFERENCES users (id) ON DELETE SET NULL,
			CONSTRAINT fk_reporter FOREIGN KEY (reporter_id) REFERENCES users (id) ON DELETE SET NULL,
			CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
		);`,

		// Create comments table
		`CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			description VARCHAR(255),
			ticket_id INTEGER NOT NULL,
			CONSTRAINT fk_ticket FOREIGN KEY (ticket_id) REFERENCES tickets (id) ON DELETE CASCADE
		);`,
	}

	for _, query := range queries {
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			log.Fatalf("Error executing query: %v\nQuery: %s", err, query)
		}
		fmt.Println("Table created or already exists.")
	}
}
