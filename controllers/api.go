package controllers

import(
	"net/http"
    "../models"
    "encoding/json"
)

func API(w http.ResponseWriter, r *http.Request) {

    var counter models.Counter

    data := json.NewDecoder(r.Body)
    err := data.Decode(&counter)

    if err != nil {
        // error
    }

    switch r.URL.Path {
        case "/api/create/":
    		counter.Create();
        case "/api/increment/":
    		counter.Increment();
            counter.Update();
        case "/api/decrement/":
    		counter.Decrement();
            counter.Update();
        case "/api/update/":
    		counter.Update();
    	default:
            // error
    }

    json.NewEncoder(w).Encode(counter)

}
