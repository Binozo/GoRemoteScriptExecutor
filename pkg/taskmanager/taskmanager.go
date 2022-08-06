package taskmanager

import (
	"RemoteScriptExecutor/pkg/system"
	"github.com/go-co-op/gocron"
	"time"
)

func Setup() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hours().Do(CheckForUpdate)
	s.StartAsync()
}

func CheckForUpdate() {
	system.CheckForUpdate()
}
