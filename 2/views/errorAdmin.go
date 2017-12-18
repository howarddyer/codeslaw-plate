package views

import(
    "net/http"
    "../config"
)

func ErrorAdmin(w http.ResponseWriter) {
    config.TemplatesAdmin.ExecuteTemplate(w, "error", nil)
}
