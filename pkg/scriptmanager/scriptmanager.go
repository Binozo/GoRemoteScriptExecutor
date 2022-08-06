package scriptmanager

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
	var files []string
	err := filepath.Walk(ScriptsDirName, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
