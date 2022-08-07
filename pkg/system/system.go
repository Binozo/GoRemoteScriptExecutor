package system

import (
	"RemoteScriptExecutor/pkg/constants"
	"context"
	"github.com/google/go-github/v45/github"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetArchitecture() string {
	return runtime.GOARCH
}

// CheckForUpdate checks if an update for GoRemoteScriptExecutor is available on the official GitHub Repository and installs it
func CheckForUpdate() {
	log.Println("Checking for updates...")
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), constants.Developer, constants.Name)
	if err != nil {
		log.Printf("Couldn't check for updates: %e. Trying again in one hour.", err)
	}
	newTagName := release.GetTagName()
	currentTagName := "v" + constants.Version

	if newTagName != currentTagName {
		log.Printf("New version %s is available...", newTagName)
		downloadUpdateAndInstall(release)
	}
}
