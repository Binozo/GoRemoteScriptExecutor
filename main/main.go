package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	password := getPassword()
	http.HandleFunc("/run", func(w http.ResponseWriter, req *http.Request) {
		scriptName := req.URL.Query().Get("script")

		result := password == req.URL.Query().Get("password")

		if !result {
			fmt.Println("Password wrong")
			w.Write([]byte("Password wrong"))
			return
		}

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

func getPassword() string {
	f, err := os.Open("creds.txt")

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return ""

}
