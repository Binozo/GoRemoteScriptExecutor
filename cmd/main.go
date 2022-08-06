package main

import (
	"RemoteScriptExecutor/pkg/arghandler"
	"RemoteScriptExecutor/pkg/credentialsmanager"
	"RemoteScriptExecutor/pkg/scriptmanager"
	"RemoteScriptExecutor/pkg/server"
	"RemoteScriptExecutor/pkg/taskmanager"
)

func main() {
	arghandler.HandleArgs() // Note: HandleArgs is capable of terminating the entire application
	taskmanager.Setup()
	credentialsmanager.CheckIfPasswordIsSet() // Note: May terminates the application if the password hasn't been set correctly
	scriptmanager.Setup()
	server.RegisterRoutes()
	server.Run()
}
