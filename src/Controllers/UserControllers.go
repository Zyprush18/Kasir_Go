package controllers

import (
	"encoding/json"

	"net/http"

	databases "github.com/Zyprush18/Kasir_Go/src/Databases"
	models "github.com/Zyprush18/Kasir_Go/src/Models"
	"github.com/Zyprush18/Kasir_Go/src/helper"
	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Errors  any    `json:"error,omitempty"`
}

func ShowAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// mengecek apakah request method nya itu adalah GET
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Request Not Allowed",
		})
		return
	}

	// membuat variabel user dengan tipe data models.User
	var user []models.User

	// mengambil data user dari database
	if err := databases.DB.Find(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Get User",
			Errors:  err.Error(),
		})
		return
	}

	// response

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success",
		Data:     user,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// mengecek apakah request method nya itu adalah POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Request Not Allowed",
		})
		return
	}

	// membuat variabel user dengan tipe data models.UserRequest
	var user *models.UserRequest

	// mengambil data dari body request dan memasukkannya ke variabel user (Body Request/Body Parser)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Invalid JSON",
			Errors:  err.Error(),
		})
		return
	}

	// validasi inputan user
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field()+" is required")
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Field harus di isi",
			Errors:  validationErrors,
		})
		return
	}

	// bcrypt password / hashing password
	hash, err := helper.HashingPassword(user.PASSWORD)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Hashing Password",
			Errors:  err.Error(),
		})
		return
	}

	// membuat variabel CreateUser dengan tipe data models.User
	CreateUser := models.User{
		USERNAME: user.USERNAME,
		PASSWORD: hash,
		EMAIL:    user.EMAIL,
	}

	// mengecek apakah ada Email yang sama
	if err := databases.DB.Where("email = ?", user.EMAIL).First(&CreateUser).Error; err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Email Already Exists",
		})
		return
	}

	// membuat data user / insert data user ke database
	if err := databases.DB.Create(&CreateUser).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed  Create User",
			Errors:  err.Error(),
		})
		return
	}

	// response
	w.WriteHeader(http.StatusCreated)
	// menampilkan response
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success Created User",
		Data:     []models.User{CreateUser},
	})
}

func ShowUserById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// mengecek apakah yang di kirim itu adalah method get
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Not Allowed",
		})
		return
	}

	// mengambil id dari url /users/{id}
	id := r.URL.Path[len("/users/"):]
	// mengecek apakah id nya ada atau tidak
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Not Found",
		})
		return
	}

	var user []models.User

	// mengecek apakah id yg di masukkan sama atau tidak
	if err := databases.DB.Where("id = ?", id).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "User not found",
		})
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "Success",
		Data:     user,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//mengecek apakah method nya itu adalah patch atau bukan 
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Not Allowed",
		})
		return
	}

	var requestUser *models.UserRequest

	// mengambil data dari body request dan memasukkannya ke variabel requestUser (Body Request/Body Parser)
	if err := json.NewDecoder(r.Body).Decode(&requestUser);err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Request body must be a valid JSON object",
		})
		return
	}

	// mengambil id dari url
	id := r.URL.Path[len("/users/update/"):]

	var user *models.User

	// melakukan pengecekkan apakah id nya ada atau tidak
	if err := databases.DB.Where("id= ?",id).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Messages{
			Messages: "User Not Found",
		})
		return
	}

	// Deklarasi map kosong dari tipe string
	userResponse := map[string]interface{}{}

	// pengecekan apakah username akan di ubah atau tidak
	if requestUser.USERNAME != "" {
		userResponse["USERNAME"] = requestUser.USERNAME
	}

	// pengecekan apakah email akan di ubah atau tidak
	if requestUser.EMAIL != "" {
		userResponse["EMAIL"] = requestUser.EMAIL
	}

	// pengecekan apakah password akan di ubah atau tidak
	if requestUser.PASSWORD != "" {
		// melakukan hashing password
		hashing,err := helper.HashingPassword(requestUser.PASSWORD)

		// jika terjadi error ketika melakukan hashing password
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorMessage{
				Message: "Failed Hashing Password",
				Errors:  err.Error(),
			})
			return
		}
		userResponse["PASSWORD"] = hashing 
	}

	// melakukan update 
	if err := databases.DB.Model(&user).Updates(userResponse).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Update User",
		})
	}

	// response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "User Updated Successfully",
	})
	
}


func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Method Not Allowed",
		})
		return 
	}

	id := r.URL.Path[len("/users/delete/"):]

	var user *models.User

	if err := databases.DB.First(&user, "id = ?",id).Delete(&user).Error;err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(ErrorMessage{
			Message: "Failed Delete User",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Messages{
		Messages: "User Delete Succcessly",
	})
}
