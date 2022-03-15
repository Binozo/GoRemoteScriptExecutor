package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/run", func(w http.ResponseWriter, req *http.Request) {
		scriptName := req.URL.Query().Get("script")

		length := len(scriptName)
		if length == 0 {
			fmt.Println("No script name provided")
			w.Write([]byte("No script name provided"))
			return
		}

		cmd := exec.Command("/bin/bash", scriptName)
		fmt.Println(cmd.Output())

		w.Write([]byte("{\"result\":\"success\"}"))

	})
	http.ListenAndServe(":5123", nil)
}
