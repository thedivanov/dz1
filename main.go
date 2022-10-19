package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/health/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Status: "OK"}
		respJSON, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(respJSON)
	})

	http.ListenAndServe(":8000", nil)
}
