package views

import(
    "net/http"
    "../config"
)

func User(w http.ResponseWriter) {
    config.TemplatesUser.ExecuteTemplate(w, "user", nil)
}
