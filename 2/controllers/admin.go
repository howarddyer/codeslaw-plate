package controllers

import(
	"net/http"
    "../views"
)

func Admin(w http.ResponseWriter, r *http.Request) {

    switch r.URL.Path {
        case "/admin/":
            views.Admin(w)
        default:
    		views.Error(w)
    }

}
