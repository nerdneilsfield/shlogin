package main

import (
	"os"

	"github.com/nerdneilsfield/shlogin/cmd"
	loggerPkg "github.com/nerdneilsfield/shlogin/pkg/logger"
)

var (
	version   = "dev"
	buildTime = "unknown"
	gitCommit = "unknown"
)

var logger = *loggerPkg.GetLogger()

func main() {
	if err := cmd.Execute(version, buildTime, gitCommit); err != nil {
		logger.Error("Failed to execute root command", "err:", err)
		os.Exit(1)
	}
}
