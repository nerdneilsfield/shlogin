package cmd

import (
	"fmt"
	"strings"

	"github.com/nerdneilsfield/shlogin/internal/configs"
	"github.com/spf13/cobra"
)

func newGenConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "gen",
		Short:        "Generate an example config file: toml or json",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			configType := "toml"
			if len(args) > 0 {
				configType = strings.ToLower(args[0])
			}
			switch configType {
			case "toml":
				return configs.GenDefaultConfigToml("./config_example.toml")
			case "json":
				return configs.GenDefaultConfigJson("./config_example.json")
			default:
				return fmt.Errorf("invalid config type: %s", configType)
			}
		},
	}
}
