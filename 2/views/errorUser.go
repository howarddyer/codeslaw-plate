package views

import(
    "net/http"
    "../config"
)

func ErrorUser(w http.ResponseWriter) {
    config.TemplatesUser.ExecuteTemplate(w, "error", nil)
}
