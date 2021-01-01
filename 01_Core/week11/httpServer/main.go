// Create a http Server
package main 

import "net/http"

func main(){
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Lilah"))
		// cast into a byte array instead of just a string
	})
	// visit localhost:8080/hello-world to view the text
	// Choose to use a JSON encoder to display JSON object/data or template package to return an html page
	http.ListenAndServe(":8080", nil)
	// Listen&serve function from the http library launches the server on the default port 8080
	// run main.go to see it running in the background - serves nothing
		
}  