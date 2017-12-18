package controllers

import(
	"net/http"
    "../views"
)

func User(w http.ResponseWriter, r *http.Request) {

    switch r.URL.Path {
        case "/user/":
            views.User(w)
        default:
    		views.ErrorUser(w)
    }

}
