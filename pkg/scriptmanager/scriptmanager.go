package scriptmanager

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	ScriptsDirName = "scripts"
)

func Setup() {
	if _, err := os.Stat(ScriptsDirName); os.IsNotExist(err) {
		// scripts directory does not exist
		// create the dir
		log.Println("The \"", ScriptsDirName, "\" directory is missing. Creating it...")
		err := os.Mkdir(ScriptsDirName, os.ModePerm)
		if err != nil {
			log.Fatal("Error: Couldn't create the", ScriptsDirName, "directory:", err, ". Exiting.")
			panic(err)
		}
	}
}

func GetScripts() ([]string, error) {
	var fileNames []string
	files, err := ioutil.ReadDir(ScriptsDirName)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames, nil
}
