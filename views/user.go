package views

import(
    "net/http"
    "../config"
)

func User(w http.ResponseWriter) {
    config.Templates.User.ExecuteTemplate(w, "user", nil)
}
