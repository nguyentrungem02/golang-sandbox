package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"info":    "Hoc golang with gin",
		"message": "Hello World with Golang!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Course", "Hoc golang")

	//data, err := json.Marshal(response)
	//if err != nil {
	//	http.Error(w, "Loi ma hoa json", http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Write(data)

	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/demo", demoHandler)

	log.Println("Server is starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Error:", err)
	}
}
