package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://chat.whatsapp.com/E75NYG8eKvyEXk6QtFtj92", http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}
