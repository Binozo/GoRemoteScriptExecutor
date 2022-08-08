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
	} else {
		log.Println("No updates available. Current version is" + currentTagName + "and newest version is" + newTagName)
	}
}

func downloadUpdateAndInstall(release *github.RepositoryRelease) {
	/* Example for asset list names:
	[
		goremotescriptexecutor_amd64
		goremotescriptexecutor_arm
		goremotescriptexecutor_arm64
	]
	*/
	target := strings.ToLower(constants.Name) + "_" + GetArchitecture()
	downloadUrl := ""
	for _, asset := range release.Assets {
		if asset.GetName() == target {
			downloadUrl = asset.GetBrowserDownloadURL()
			break
		}
	}
	if downloadUrl == "" {
		log.Printf("Couldn't find executable for architecture %s. Update aborted. Please compile it yourself or create an issue and demand a download for your architecture.", GetArchitecture())
		return
	}
	log.Printf("Downloading update from %s", downloadUrl)
	executableName := filepath.Base(os.Args[0])
	fileName := executableName + ".new"
	err := downloadFile(fileName, downloadUrl)
	if err != nil {
		log.Printf("Couldn't download update: %s. Retrying in one hour", err.Error())
		return
	}
	log.Println("Update downloaded. Testing executable..")
	err = os.Chmod(fileName, 0755)
	fullPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Printf("Couldn't get absolute path of executable: %s. Aborting update.", err.Error())
		return
	}
	response, err := exec.Command(fullPath, constants.AliveArg).Output()
	if err != nil {
		log.Printf("Couldn't test executable: %s. Maybe it's corrupt or doesn't work correctly on your platform. Aborting update.", err.Error())
		os.Remove(fileName)
		return
	}
	if string(response) != constants.AliveResponse {
		log.Printf("Executable doesn't work. Aborting update.")
		os.Remove(fileName)
		return
	}
	log.Println("Executable works. Installing update...")
	err = os.Remove(executableName)
	if err != nil {
		log.Printf("Couldn't remove old executable: %s", err.Error())
	}
	err = os.Rename(fileName, executableName)
	if err != nil {
		log.Printf("Couldn't rename new executable: %s. Aborting update.", err.Error())
		return
	}
	log.Println("Update installed. Restarting...")
	os.Exit(2)
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
