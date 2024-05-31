package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleform(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { //parse errors err:= r.ParseForm(); if err != nil { return }
		fmt.Fprintf(w, "Parseform() err %v", err)
		return
	}

	fmt.Fprintf(w, "ParseForm Sucess!!")

	name := r.FormValue("name") //r.FormVallue get data
	addr := r.FormValue("email")

	fmt.Fprintf(w, "Name : %v", name) //reponse
	fmt.Fprintf(w, "Email : %v", addr)
}

func handlehello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found!!", http.StatusNotFound) //check if path same r.URL.Path
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Found", http.StatusBadRequest) //check method with r.Method
		return
	}

	fmt.Fprintf(w, "Hello Arnav!") //Fprintf(w, String) for response
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //assigning static files to doc

	http.Handle("/", fileServer)           //allowing a handler to serve the request
	http.HandleFunc("/form", handleform)   //allowing a function to handle request
	http.HandleFunc("/hello", handlehello) //allowing function to handle

	fmt.Printf("Server started at port 8080 \n") //Printf same as printf

	if err := http.ListenAndServe(":8080", nil); err != nil { //check port working
		log.Fatal(err)
	}
}
