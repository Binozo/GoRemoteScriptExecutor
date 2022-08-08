package server

import (
	"RemoteScriptExecutor/pkg/constants"
	"RemoteScriptExecutor/pkg/credentialsmanager"
	"RemoteScriptExecutor/pkg/scriptmanager"
	"RemoteScriptExecutor/pkg/system"
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
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/scripts", scriptsHandler).Methods("GET")
	r.HandleFunc("/runScript/{script}", runScriptHandler).Methods("GET")
	r.HandleFunc("/update", updateHandler).Methods("GET")
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

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if res := checkCredentials(w, r); res == false {
		return
	}
	system.CheckForUpdate()

	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(map[string]interface{}{
		"status":  "ok",
		"message": "Check for updates has been started",
	})
	w.Write(json)
}

func runScriptHandler(w http.ResponseWriter, r *http.Request) {
	if res := checkCredentials(w, r); res == false {
		return
	}
	// Check if script name is in request
	scriptName := mux.Vars(r)["script"]
	blocking := r.URL.Query().Get("blocking") == "true"
	responseOutput := r.URL.Query().Get("responseOutput") == "true"
	if scriptName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(map[string]string{
			"status":  "error",
			"message": "No script name provided",
		})
		w.Write(json)
		return
	}

	// Check if the script exists
	scripts, err := scriptmanager.GetScripts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		w.Write(json)
	}

	scriptExists := false
	for _, curScriptName := range scripts {
		if curScriptName == scriptName {
			scriptExists = true
		}
	}
	if !scriptExists {
		w.WriteHeader(http.StatusBadRequest)
		json, _ := json.Marshal(map[string]string{
			"status":  "error",
			"message": "Script doesn't exist",
		})
		w.Write(json)
		return
	}

	// Run the script
	if !blocking {
		go func() {
			_, err := scriptmanager.ExecuteScript(scriptName)
			if err != nil {
				log.Println("Error while executing script:", err)
			}
		}()
		w.WriteHeader(http.StatusOK)
		json, _ := json.Marshal(map[string]string{
			"status": "ok",
			"mode":   "non-blocking",
		})
		w.Write(json)
		return
	}
	output, err := scriptmanager.ExecuteScript(scriptName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal(map[string]string{
			"status":  "error",
			"type":    "Error while executing script",
			"message": err.Error(),
		})
		w.Write(json)
		return
	}
	if responseOutput {
		w.WriteHeader(http.StatusOK)
		json, _ := json.Marshal(map[string]string{
			"status": "ok",
			"output": output,
		})
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusOK)
		json, _ := json.Marshal(map[string]string{
			"status": "ok",
		})
		w.Write(json)
	}
}
