package views

import(
    "net/http"
    "../config"
)

func Public(w http.ResponseWriter) {
    config.TemplatesPublic.ExecuteTemplate(w, "public", nil)
}
