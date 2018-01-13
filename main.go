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
	http.HandleFunc("/api/", controllers.API)

    // Listen on port
	http.ListenAndServe(config.Ports.Test, nil)

}
