package credentialsmanager

import (
	"RemoteScriptExecutor/pkg/constants"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	return hashedPassword
}

func GeneratePassword(password string) error {
	return os.Setenv(constants.EnvPassword, HashPassword(password))
}

func CheckIfPasswordIsSet() {
	value, existing := os.LookupEnv(constants.EnvPassword)
	if !existing {
		log.Fatalln("Password hasn't been set. https://github.com/Binozo/GoRemoteScriptExecutor#Setup")
		os.Exit(1)
	}
	if value == "" {
		log.Fatalln("Please set a good password. https://github.com/Binozo/GoRemoteScriptExecutor#Setup")
		os.Exit(1)
	}
}

func GetPassword() string {
	envPassword := os.Getenv(constants.EnvPassword)
	if envPassword == "" {
		log.Fatalln("Please set a good password. https://github.com/Binozo/GoRemoteScriptExecutor#Setup")
		os.Exit(1)
	}
	return envPassword
}
