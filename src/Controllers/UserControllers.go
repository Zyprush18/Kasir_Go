package controllers

import (
	"encoding/json"

	"net/http"

	databases "github.com/Zyprush18/Kasir_Go.git/src/Databases"
	models "github.com/Zyprush18/Kasir_Go.git/src/Models"
	"github.com/Zyprush18/Kasir_Go.git/src/helper"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Errors any `json:"error"`
}

func ShowAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// mengecek apakah request method nya itu adalah GET
	if r.Method != http.MethodGet {
		w.WriteHeader( http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Request Not Allowed",
		})
		return
	}

	// membuat variabel user dengan tipe data models.User
	var user []models.User

	// mengambil data user dari database
	if err := databases.DB.Find(&user).Error; err != nil{
		w.WriteHeader( http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Get User",
			Errors: err.Error(),
		})
		return
	}

	// response
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success",
		Data: user,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	// mengecek apakah request method nya itu adalah POST
	if r.Method != http.MethodPost {
		w.WriteHeader( http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Request Not Allowed",
		})
		return
	}

	// membuat variabel user dengan tipe data models.UserRequest
	var user *models.UserRequest

	// mengambil data dari body request dan memasukkannya ke variabel user (Body Request/Body Parser)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader( http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Invalid JSON",
			Errors: err.Error(),
		})
		return
	}

	// validasi inputan user
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field() + " is required")
		}

		w.WriteHeader( http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Field harus di isi",
			Errors: validationErrors,
		})
		return
	}

	// bcrypt password / hashing password
	hash,err := helper.HashingPassword(user.PASSWORD)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Hashing Password",
			Errors: err.Error(),
		})
		return
	}

	// membuat variabel CreateUser dengan tipe data models.User
	CreateUser := models.User{
		USERNAME: user.USERNAME,
		PASSWORD: hash,
		EMAIL: user.EMAIL,
	}

	// mengecek apakah ada Email yang sama
	if err := databases.DB.Where("email = ?", user.EMAIL).First(&CreateUser).Error;err == nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Email Already Exists",
		})
		return
	}


	// membuat data user / insert data user ke database
	if err := databases.DB.Create(&CreateUser).Error; err != nil {
		w.WriteHeader( http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed  Create User",
			Errors: err.Error(),
		})
		return
	}

	// response
	w.WriteHeader(http.StatusCreated)
	// menampilkan response
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success Created User",
		Data: []models.User{CreateUser},
	})
}