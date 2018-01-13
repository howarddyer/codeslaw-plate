package views

import(
    "net/http"
    "../config"
    "../models"
)

func Public(w http.ResponseWriter) {
    page := models.Page{
        Title: "Home",
        Description: "Welcome to the home page",
        Type: "Public",
    }

    page.UpdateContent("Welcome to our site")

    config.Templates.Public.ExecuteTemplate(w, "public", page)
}
