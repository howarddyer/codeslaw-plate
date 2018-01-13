package main

import(
	"net/http"
    "./controllers"
    "./config"
    "./models"
)

func main() {

    // Initialise database
    models.InitDatabase("./main.db")

    // Handle root URL
    http.HandleFunc("/", controllers.Public)

    // Handle non-root URL
	http.HandleFunc("/admin/", controllers.Admin)
	http.HandleFunc("/user/", controllers.User)

    // Listen on port
	http.ListenAndServe(config.Ports.Test, nil)

}
