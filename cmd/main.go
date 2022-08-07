package main

import (
	"RemoteScriptExecutor/pkg/arghandler"
	"RemoteScriptExecutor/pkg/constants"
	"RemoteScriptExecutor/pkg/credentialsmanager"
	"RemoteScriptExecutor/pkg/scriptmanager"
	"RemoteScriptExecutor/pkg/server"
	"RemoteScriptExecutor/pkg/system"
	"RemoteScriptExecutor/pkg/taskmanager"
	"fmt"
)

func main() {
	arghandler.HandleArgs() // Note: HandleArgs is capable of terminating the entire application
	fmt.Printf("Starting %s v%s (%s Edition)\n", constants.Name, constants.Version, system.GetArchitecture())
	taskmanager.Setup()
	credentialsmanager.CheckIfPasswordIsSet() // Note: May terminates the application if the password hasn't been set correctly
	scriptmanager.Setup()
	server.RegisterRoutes()
	server.Run()
}
