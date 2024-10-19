package cmd

import (
	"fmt"

	"github.com/inconshreveable/log15"
	loggerPkg "github.com/nerdneilsfield/shlogin/pkg/logger"
	"github.com/spf13/cobra"
)

var verbose bool

func newRootCmd(version string, buildTime string, gitCommit string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shlogin",
		Short: "shlogin is a tool to login to the ShanghaiTech network",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				loggerPkg.UpdateLogLevel(log15.LvlDebug)
			} else {
				loggerPkg.UpdateLogLevel(log15.LvlInfo)
			}
		},
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	cmd.AddCommand(newVersionCmd(version, buildTime, gitCommit))
	cmd.AddCommand(newGenConfigCmd())
	cmd.AddCommand(newCheckConfigCmd())
	cmd.AddCommand(newConvertConfigCmd())
	cmd.AddCommand(newEditConfigCmd())
	cmd.AddCommand(newConnectionCmd())
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newCronCmd())
	return cmd
}

func Execute(version string, buildTime string, gitCommit string) error {
	if err := newRootCmd(version, buildTime, gitCommit).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
