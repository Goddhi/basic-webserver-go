package main

import (
	"fmt"
	"log"
	"net/http"
)

func forwardHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {  // if there is an while trying to pass the form
		fmt.Fprintf(w, "ParseForm() err: %v", err)  // print out the error
		return
	}
	fmt.Fprintf(w, "POST request successful \n") // post the string to the screen
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name) // prints the value of name  to the screen
	fmt.Fprintf(w, "Address = %s\n", address) /// prints the value of addresss to the screen
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { // if another url path is requested aside from the hello path 
		http.Error(w, "404 not founbd", http.StatusNotFound)
		return 
	}

	if r.Method != "GET" {  // if method is not GET, print an error
		http.Error(w, "Method is not supported ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello GOddhi") // output of the /home path are printed to the screen
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // create a fileServer from static directory
	http.Handle("/", fileServer) // creating home route
	http.HandleFunc("/form", forwardHandler) // create the form route 
	http.HandleFunc("/hello", helloHandler) /// create the home route

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // creating a webserver that listen on port 8080 and catch errors. 
		log.Fatal(err) // if there is an error print out the error
	}

}

// commands to tun the code
// go mod init go-server
// go build
// go run main.go


