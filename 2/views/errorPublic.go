package views

import(
    "net/http"
    "../config"
)

func ErrorPublic(w http.ResponseWriter) {
    config.Templates.Public.ExecuteTemplate(w, "error", nil)
}
