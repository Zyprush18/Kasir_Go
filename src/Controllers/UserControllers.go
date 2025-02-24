package controllers

import (
	"encoding/json"
	"net/http"

	databases "github.com/Zyprush18/Kasir_Go.git/src/Databases"
	models "github.com/Zyprush18/Kasir_Go.git/src/Models"
	"github.com/go-playground/validator/v10"
)



func ShowAllUser(w http.ResponseWriter, r *http.Request) {
	// mengecek apakah request method nya itu adalah GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// membuat variabel user dengan tipe data models.User
	var user []models.User

	// mengambil data user dari database
	if err := databases.DB.Find(&user).Error; err != nil{
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success",
		Data: user,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	// mengecek apakah request method nya itu adalah POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// membuat variabel user dengan tipe data models.UserRequest
	var user models.UserRequest

	// mengambil data dari body request dan memasukkannya ke variabel user (Body Request/Body Parser)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w,"Not Field Created", http.StatusInternalServerError)
		return
	}

	// validasi inputan user
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// membuat variabel CreateUser dengan tipe data models.User
	CreateUser := models.User{
		USERNAME: user.USERNAME,
		PASSWORD: user.PASSWORD,
		EMAIL: user.EMAIL,
	}

	// membuat data user / insert data user ke database
	if err := databases.DB.Create(&CreateUser).Error; err != nil {
		http.Error(w, "Failed Create User", http.StatusBadRequest)
		return
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// menampilkan response
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success Created User",
		Data: []models.User{CreateUser},
	})
}