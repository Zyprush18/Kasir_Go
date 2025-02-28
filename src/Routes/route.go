package routes

import (
	"fmt"
	"net/http"

	controllers "github.com/Zyprush18/Kasir_Go/src/Controllers"
)

func Route() {
	// user
	http.HandleFunc("/users", controllers.ShowAllUser)
	http.HandleFunc("/users/created", controllers.CreateUser)
	http.HandleFunc("/users/{id}", controllers.ShowUserById)
	http.HandleFunc("/users/update/{id}", controllers.UpdateUser)
	http.HandleFunc("/users/delete/{id}", controllers.DeleteUser)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
