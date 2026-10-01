package main

import (
	"fmt"
	"log"
	"net/http"
)


func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"world, hello!")
}

func handlerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from different page...")

	if r.Method == http.MethodGet {
			fmt.Fprintln(w, "youd did get")
	}else if r.Method == http.MethodPost{
			fmt.Fprintln(w, "you did post")
	} else {
			fmt.Fprintln(w, "error: non")
	}



}


func main() {
	
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/about", handlerPage)
	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))


	/*
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/about", handlerPage)


	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	*/

}
