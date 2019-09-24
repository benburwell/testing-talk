package main

import (
	"io/ioutil"
	"net/http"

	"bnbl.io/caesar"
)

func main() {
	server := handleCaesar()
	http.ListenAndServe(":8088", server)
}

func handleCaesar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		encoded := caesar.Encode(string(body))
		w.Write([]byte(encoded))
	}
}
