package views

import(
    "net/http"
    "../config"
)

func ErrorUser(w http.ResponseWriter) {
    config.Templates.User.ExecuteTemplate(w, "error", nil)
}
