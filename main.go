package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Kullanıcıları JSON olarak döndüren API endpoint
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := getUsers() // users.go'dan getUsers fonksiyonunu çağırıyoruz
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	// JSON formatında kullanıcıları yanıt olarak gönder
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Kullanıcı ekleme API endpoint
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// POST isteğinden gelen JSON verisini okuma
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Yeni kullanıcıyı veritabanına ekleme
	newUser, err := addUser(user.Username, user.Email) // users.go'dan addUser fonksiyonunu çağırıyoruz
	if err != nil {
		http.Error(w, "Failed to add user", http.StatusInternalServerError)
		return
	}

	// JSON formatında eklenen kullanıcıyı yanıt olarak gönder
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users", getUsersHandler)    // Kullanıcıları al
	http.HandleFunc("/users/add", addUserHandler) // Kullanıcı ekle

	// Sunucuyu başlat
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
