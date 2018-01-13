package views

import(
    "net/http"
    "../config"
)

func Admin(w http.ResponseWriter) {
    config.Templates.Admin.ExecuteTemplate(w, "admin", nil)
}
