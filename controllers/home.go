package controllers

import(
  "net/http"
  "library"
)

func Home(w http.ResponseWriter, r *http.Request) {
  err := library.Templates.ExecuteTemplate(w, "home", nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
