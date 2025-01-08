package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// User yapısı
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Veritabanına bağlanma fonksiyonu
func connectToDB() (*sql.DB, error) {
	connStr := "postgres://postgres:password@localhost:5433/go_test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Kullanıcıyı veritabanına eklemek için fonksiyon
func addUser(username, email string) (*User, error) {
	db, err := connectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Kullanıcıyı ekleyen SQL sorgusu
	var user User
	err = db.QueryRow("INSERT INTO users(username, email) VALUES($1, $2) RETURNING id", username, email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	// Kullanıcıyı döndür
	user.Username = username
	user.Email = email
	return &user, nil
}

// Veritabanından tüm kullanıcıları getiren fonksiyon
func getUsers() ([]User, error) {
	db, err := connectToDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Hata kontrolü
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
