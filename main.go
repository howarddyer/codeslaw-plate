package main

import(
  "net/http"
  "controllers"
)

func init() {
  http.HandleFunc("/", controllers.Home)
}
