package main

import(
	"net/http"
    "./controllers"
)

func main() {

	http.HandleFunc("/", controllers.Public)
	http.HandleFunc("/user/", controllers.User)
    http.HandleFunc("/admin/", controllers.Admin)

	http.ListenAndServe(":8000", nil)

}
