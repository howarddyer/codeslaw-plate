package controllers

import(
	"net/http"
    "../views"
)

func Public(w http.ResponseWriter, r *http.Request) {

    switch r.URL.Path {
    	case "/":
    		views.Home(w)
    	default:
    		views.Error(w)
    }

}
