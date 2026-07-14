package main

import (

)

func main(){

	mux := http.NewServeMux() //Creating router

	//Registering handlers
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artists/", handlers.ArtistHandler)

	//Telling the system where our static files are
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static",fs)),

		fmt.Println("Server running ....")

	//Starting the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
		http.Errorf("InternalServerError")
	}
}

//main here basically, create server router (mux), tells where to
//locate static files, called handlers and starts the server.