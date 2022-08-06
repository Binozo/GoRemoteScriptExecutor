package arghandler

import (
	"RemoteScriptExecutor/pkg/constants"
	"RemoteScriptExecutor/pkg/credentialsmanager"
	"fmt"
	"log"
	"os"
)

func HandleArgs() {
	if len(os.Args) <= 1 {
		log.Println("No arguments have been passed")
		return
	}

	if os.Args[1] == constants.AliveArg {
		fmt.Print(constants.AliveResponse)
		os.Exit(0)
	} else if len(os.Args) >= 3 && os.Args[1] == constants.SetupPasswordArg && os.Args[2] != "" {
		log.Println("Setting password...")
		if err := credentialsmanager.GeneratePassword(os.Args[2]); err != nil {
			log.Fatalf("Couldn't set the password env variable: %e", err)
			os.Exit(1)
		} else {
			log.Println("Password successfully set. Exiting")
			os.Exit(0)
		}
	}
}
