package credentialsmanager

import (
	"RemoteScriptExecutor/pkg/constants"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func GeneratePassword(password string) error {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	return os.Setenv(constants.EnvPassword, hashedPassword)
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
