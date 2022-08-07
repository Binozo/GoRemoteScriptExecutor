package server

import (
	"RemoteScriptExecutor/pkg/constants"
	"RemoteScriptExecutor/pkg/credentialsmanager"
	"RemoteScriptExecutor/pkg/scriptmanager"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	r *mux.Router
)

func RegisterRoutes() {
	r = mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/scripts", scriptsHandler)
	http.Handle("/", r)
}

func Run() {
	log.Println("Listening on Port", constants.Port, "...")
	err := http.ListenAndServe(":"+strconv.Itoa(constants.Port), nil)
	log.Fatal("HTTP Error:", err, ". Restarting...")
	Run()
}

func checkCredentials(w http.ResponseWriter, r *http.Request) bool {
	clientAuthKey := r.Header.Get("Authorization")
	_, hashedClientAuthKey := credentialsmanager.HashPassword(clientAuthKey)
	// Now the AuthKey needs to be hashed and compared to the one stored in the environment variable
	if !bytes.Equal(hashedClientAuthKey, credentialsmanager.GetPassword()) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return false
	}
	return true
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	jsonRes, _ := json.Marshal(map[string]string{
		"name":    constants.Name,
		"version": constants.Version,
	})
	w.Write(jsonRes)
}

func scriptsHandler(w http.ResponseWriter, r *http.Request) {
	if res := checkCredentials(w, r); res == false {
		return
	}
	scripts, err := scriptmanager.GetScripts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		w.Write(json)
	}
	if len(scripts) == 0 {
		// bugfix: if the slice is empty, json.Marshal returns "null" instead of "[]" => https://imantung.medium.com/golang-json-marshal-return-null-for-empty-slice-9aa816b7324b
		scripts = make([]string, 0)
	}

	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(map[string]interface{}{
		"status":  "ok",
		"scripts": scripts,
	})
	w.Write(json)
}
