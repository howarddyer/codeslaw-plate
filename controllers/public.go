package controllers

import(
	"net/http"
    "../views"
)

func Public(w http.ResponseWriter, r *http.Request) {

    switch r.URL.Path {
        case "/":
    		views.Public(w)
    	default:
    		views.ErrorPublic(w)
    }

}
