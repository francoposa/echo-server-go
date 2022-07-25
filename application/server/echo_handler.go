package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
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

	w.WriteHeader(http.StatusOK)
}
