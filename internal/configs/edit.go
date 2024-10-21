package configs

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"go.uber.org/zap"
)

func EditConfigFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		logger.Error("Config file not found", zap.String("path", path))
		return err
	}

	if !strings.HasSuffix(path, ".toml") && !strings.HasSuffix(path, ".json") {
		logger.Error("Unsupported config file format", zap.String("path", path))
		return fmt.Errorf("unsupported config file format: %s", path)
	}

	var editCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		logger.Debug("Opening config file in Windows", zap.String("path", path))
		editCmd = exec.Command("cmd", "-c", "start", path)
	case "darwin":
		logger.Debug("Opening config file in macOS", zap.String("path", path))
		editCmd = exec.Command("open", "-a", "TextEdit", path)
	case "linux":
		logger.Debug("Opening config file in Linux", zap.String("path", path))
		editCmd = exec.Command("xdg-open", path)
	default:
		logger.Debug("Opening config file in default editor", zap.String("path", path))
		editCmd = exec.Command("vim", path)
	}

	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	err := editCmd.Run()
	if err != nil {
		logger.Error("Failed to open config file in editor", zap.String("path", path), zap.Error(err))
		return err
	}

	return nil
}
