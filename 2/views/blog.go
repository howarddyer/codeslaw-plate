package views

import(
    "html/template"
    "net/http"
)

func Blog(w http.ResponseWriter) {
    tmpl, _ := template.ParseFiles("./templates/layout.html")
    tmpl.Execute(w, "Blog")
}
