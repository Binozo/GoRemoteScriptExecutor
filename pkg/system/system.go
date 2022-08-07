package system

import (
	"log"
	"runtime"
)

func GetArchitecture() string {
	return runtime.GOARCH
}

// CheckForUpdate checks if an update for GoRemoteScriptExecutor is available on the official GitHub Repository and installs it
func CheckForUpdate() {
	log.Println("Checking for updates...")
}
