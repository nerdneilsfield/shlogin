package cmd

import (
	"github.com/nerdneilsfield/shlogin/internal/configs"
	"github.com/nerdneilsfield/shlogin/internal/login"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "login",
		Short:        "Use config file to login to shlogin",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := configs.CheckConfig(args[0])
			if err != nil {
				logger.Error("Failed to check config", zap.Error(err))
				return err
			}
			config, err := configs.LoadConfig(args[0])
			if err != nil {
				logger.Error("Failed to load config", zap.Error(err))
				return err
			}
			return login.LoginWithConfig(config)
		},
	}
}
