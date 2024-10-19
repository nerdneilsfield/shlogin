package cmd

import (
	"github.com/nerdneilsfield/shlogin/internal/configs"
	"github.com/nerdneilsfield/shlogin/internal/login"
	loggerPkg "github.com/nerdneilsfield/shlogin/pkg/logger"
	"github.com/spf13/cobra"
)

var logger = *loggerPkg.GetLogger()

func newLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "login",
		Short:        "Use config file to login to shlogin",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := configs.CheckConfig(args[0])
			if err != nil {
				logger.Error("Failed to check config", "error", err)
				return err
			}
			config, err := configs.LoadConfig(args[0])
			if err != nil {
				logger.Error("Failed to load config", "error", err)
				return err
			}
			return login.LoginWithConfig(config)
		},
	}
}
