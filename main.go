package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Connection string for PostgreSQL
const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	dbname   = "go_test"
)

/* INSERT INTO users (username, email, created_at) VALUES
('john_doe', 'john@example.com', NOW()),
('jane_doe', 'jane@example.com', NOW()),
('mike_smith', 'mike@example.com', NOW()),
('alice_williams', 'alice@example.com', NOW()),
('bob_johnson', 'bob@example.com', NOW()),
('charlie_brown', 'charlie@example.com', NOW()),
('david_jones', 'david@example.com', NOW()),
('emily_davis', 'emily@example.com', NOW()),
('lucas_martin', 'lucas@example.com', NOW()),
('olivia_lee', 'olivia@example.com', NOW()); */

func main() {
	// Create a new router
	r := gin.Default()

	// Connect to PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}
	log.Println("Successfully connected to the database")

	// Define a simple route
	r.GET("/users", func(c *gin.Context) {
		var users []string
		rows, err := db.Query("SELECT username FROM users")
		if err != nil {
			log.Fatalf("Unable to execute query: %v\n", err)
		}
		defer rows.Close()

		for rows.Next() {
			var username string
			err = rows.Scan(&username)
			if err != nil {
				log.Fatalf("Unable to scan row: %v\n", err)
			}
			users = append(users, username)
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	})

	// Run the server
	r.Run(":8080")
}
