package cmd

import (
	"log"
	"os"
	"quiteline/ui/tui"
)

// Setup a global logging file for debugging
func setupLogger() *os.File {
	logPath := "cmd/quiteline.log"
	if err := os.Remove(logPath); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	return f
}

func Start() {
	logFile := setupLogger()
	defer logFile.Close()
	tui.Run()
}
