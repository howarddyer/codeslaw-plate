package views

import(
    "html/template"
    "net/http"
)

func Error(w http.ResponseWriter) {
    tmpl, _ := template.ParseFiles("./templates/layout.html")
    tmpl.Execute(w, "Page not found")
}
