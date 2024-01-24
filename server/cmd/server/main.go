package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/", Echo)
	mux.HandleFunc("/json", EchoJSON)

	err := http.ListenAndServe("0.0.0.0:8080", mux)
	log.Fatal(err)
}

func Echo(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func EchoJSON(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := map[string]interface{}{}
	body["method"] = r.Method
	body["protocol"] = r.Proto
	body["headers"] = r.Header
	body["remote_address"] = r.RemoteAddr
	body["body"] = string(requestBody)

	prettyJSONBody, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = w.Write(prettyJSONBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(map[string]string{"status": "ok"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
