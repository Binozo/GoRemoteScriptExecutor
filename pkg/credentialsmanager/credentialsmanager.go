package credentialsmanager

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

const (
	PasswordFileName = "pswd"
)

func HashPassword(password string) (string, []byte) {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	return hashedPassword, h.Sum(nil)
}

func GeneratePassword(password string) error {
	_, hashedPassword := HashPassword(password)
	err := os.WriteFile(PasswordFileName, hashedPassword, 0644)
	return err
}

func CheckIfPasswordIsSet() {
	if _, err := os.Stat(PasswordFileName); os.IsNotExist(err) {
		// password file does not exist
		log.Fatalln("Password hasn't been set. https://github.com/Binozo/GoRemoteScriptExecutor#Setup")
		os.Exit(1)
	}
}

func GetPassword() string {
	CheckIfPasswordIsSet()
	dat, err := os.ReadFile(PasswordFileName)
	if err != nil {
		log.Fatalln("Couldn't read the password file:", err)
		os.Exit(1)
	}
	return string(dat)
}
