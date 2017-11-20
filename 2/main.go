package main

/*
ToDo:
1. Read this site which explains the below code
   https://gowebexamples.com/hello-world/

2. Turn this simple example into a MVC structure

   controller
     |
     + opens view
        |
        + gets data using models
        |
        + loads template with the model paramters
        |
        + writes to screen
*/

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
