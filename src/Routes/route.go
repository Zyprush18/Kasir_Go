package routes

import (
	"fmt"
	"net/http"

	"github.com/Zyprush18/Kasir_Go.git/src/Controllers"
)

func Route() {
	// user
	http.HandleFunc("/user", controllers.ShowAllUser)
	http.HandleFunc("/user/created", controllers.CreateUser)








	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}