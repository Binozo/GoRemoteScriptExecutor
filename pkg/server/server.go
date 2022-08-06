package server

import (
	"RemoteScriptExecutor/pkg/constants"
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

func checkCredentials(w http.ResponseWriter, r *http.Request) {
	clientAuthKey := r.Header.Get("Authorization")

	w.WriteHeader(http.StatusOK)
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
	checkCredentials(w, r)
}
