package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"html/template"
)

//Create struct that holds the information that we want displayed on the web site
type Welcome struct {
	Name string
	Time string
}

//Go application entry point
func main() {
	//Instantiate a Welcome struct object and pass random information.

	//Get the name of the user as query parameter from URL
	welcome := Welcome{"Eric", time.Now().Format(time.Stamp)}

	//Call the html template, wrap it in the Must method for error handling.
	templates := template.Must(template.ParseFiles("welcome-template.html"))

	//HTML comes with css as well so we add the "static" directory as a url
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	//Method to take in the URL path and a function that takes a response writer and http request.
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("/" , func(ResponseWriter, *Request)) /* {
		//log access
		log.Printf("Received request for path: %s", r.URL.Path);

		//Takes the name from the URL query will set welcome.Name
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name;
		}

		//conditional for error handling. pass welcome struct to the welcome-template.html
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start web server set the port with no path as it will just be testing on localhost
	//Print any errors from starting the web server
	fmt.Println("Listening");
	fmt.Println(http.ListenAndServe(":8080", nil));
}
