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
	body["Method"] = r.Method
	body["Protocol"] = r.Proto
	body["Headers"] = r.Header
	body["RemoteAddress"] = r.RemoteAddr
	body["Body"] = string(requestBody)

	prettyJSONBody, _ := json.MarshalIndent(body, "", "    ")
	_, err = w.Write(prettyJSONBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
