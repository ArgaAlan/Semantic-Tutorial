package main

import (
    "net/http"
)

func mostrarHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "vue01.html")

}

func main() {
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}