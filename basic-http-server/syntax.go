package main

import (
	"fmt"
	"net/http"
)

// Using server's ListenAndServe function
func usingListenAndServeFunction() {
	fmt.Println("Using Listen And Serve Function")
	//Create the default mux
	mux := http.NewServeMux()

	//Handling the /v1/teachers. The handler is function here
	mux.HandleFunc("/v1/teachers", teacherHandler)

	//Handling the /v1/students. The handler is a type implementing the Handler interface here
	sHandler := studentHandler{}
	mux.Handle("/v1/students", sHandler)

	//Create the server.
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}

// Using http's ListenAndServe function
func usingHttpListenAndServeFunction() {
	fmt.Println("\nUsing Http's Listen and Serve function")

	//Handing the /v1/teachers
	http.HandleFunc("/v1/teachers", teacherHandler)

	//Handling the /v1/students
	sHandler := studentHandler{}
	http.Handle("/v1/students", sHandler)

	http.ListenAndServe(":8080", nil)
}

//Note::
// Internally both of them are doing the same thing.
// The second one uses default for everything while the first one lets you create mux and server instance explicitly so that you can specify more options and hence first option is more flexible.
