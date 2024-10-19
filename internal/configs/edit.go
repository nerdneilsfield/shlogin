package configs

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func EditConfigFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		logger.Error("Config file not found", "path", path)
		return err
	}

	if !strings.HasSuffix(path, ".toml") && !strings.HasSuffix(path, ".json") {
		logger.Error("Unsupported config file format", "path", path)
		return fmt.Errorf("unsupported config file format: %s", path)
	}

	var editCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		editCmd = exec.Command("notepad.exe", path)
	} else {
		if editor := os.Getenv("EDITOR"); editor != "" {
			editCmd = exec.Command(editor, path)
		} else {
			editCmd = exec.Command("vim", path)
		}
	}

	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	err := editCmd.Run()
	if err != nil {
		logger.Error("Failed to open config file in editor", "path", path, "err", err)
		return err
	}

	return nil
}
