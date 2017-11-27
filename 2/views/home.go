package views

import(
    "html/template"
    "net/http"
)

func Home(w http.ResponseWriter) {
    tmpl, _ := template.ParseFiles("./templates/layout.html")
    tmpl.Execute(w, "Home")
}
