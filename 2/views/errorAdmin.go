package views

import(
    "net/http"
    "../config"
)

func ErrorAdmin(w http.ResponseWriter) {
    config.Templates.Admin.ExecuteTemplate(w, "error", nil)
}
