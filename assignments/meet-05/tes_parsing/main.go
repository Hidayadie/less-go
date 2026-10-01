package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"encoding/json"
)

type UserRequest struct {
	ID int				`json: "id"`
	Name string		`json: "name"`
	Email string 	`json: "email"`
}


func HandleJSON(w http.ResponseWriter, r *http.Request) {
	// 1. Baca data mentah menjadi byte
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Gagal membaca body", http.StatusBadRequest)
		return
  }

	// 2. Cetak string mentah JSON ke terminal
	fmt.Println("--- Mentah JSON String ---")
	fmt.Println(string(bodyBytes))

	// 3. (Opsional) Mengubah ke Map untuk mengambil value tertentu
	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err == nil {
		fmt.Println("--- Hasil Parse Map ---")
		fmt.Printf("Nama: %v, Umur: %v\n", data["nama"], data["umur"])
	}

	w.Write([]byte("JSON berhasil dicetak"))
}




func handlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"world, hello!")
}

func handlerPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello from different page...")

	if r.Method == http.MethodGet {
			fmt.Fprintln(w, "youd did get")


	}else if r.Method == http.MethodPost{
			fmt.Fprintln(w, "you did post")
			var userRequest UserRequest

			err := json.NewDecoder(r.Body).Decode(&userRequest)
			if err != nil {
				fmt.Fprintln(w, "gagal parsing json")	
			return
			}

			fmt.Fprintf(w, "data json diterima %+v\n", userRequest)


	} else if r.Method == http.MethodPut{

		user_id := r.URL.Query().Get("user_id")
		fmt.Fprintln(w, "user_id yang update: ", user_id)

	} else if r.Method == http.MethodDelete{

		user_id := r.URL.Query().Get("user_id")
		fmt.Fprintln(w, "user_id yang dihapus: ", user_id)

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
