package cmd

import (
	"github.com/nerdneilsfield/shlogin/internal/configs"
	"github.com/nerdneilsfield/shlogin/internal/login"
	"github.com/spf13/cobra"
)

func newCronCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "cron",
		Short:        "Cron job to keep the connection alive",
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
			return login.CronLogin(config)
		},
	}
}
