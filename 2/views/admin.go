package views

import(
    "net/http"
    "../config"
)

func Admin(w http.ResponseWriter) {
    config.TemplatesAdmin.ExecuteTemplate(w, "admin", nil)
}
