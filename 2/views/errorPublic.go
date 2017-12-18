package views

import(
    "net/http"
    "../config"
)

func ErrorPublic(w http.ResponseWriter) {
    config.TemplatesPublic.ExecuteTemplate(w, "error", nil)
}
