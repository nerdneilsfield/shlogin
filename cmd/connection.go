package cmd

import (
	"fmt"

	"github.com/nerdneilsfield/shlogin/pkg/network"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newConnectionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "conn",
		Short:        "Check the connection",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := network.CheckWanConnection()
			if err != nil {
				return err
			}
			fmt.Println("Wan connection check success")
			return nil
		},
	}

	cmd.AddCommand(newPingCmd())
	cmd.AddCommand(newTcpPingCmd())
	cmd.AddCommand(newHttpConnectCmd())
	cmd.AddCommand(newCheckLoginServerCmd())
	cmd.AddCommand(newCheckShLanCmd())

	return cmd
}

func newPingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ping",
		Short:        "Ping the host",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			pingResult, err := network.Ping(args[0])
			if err != nil {
				return err
			}
			fmt.Println(pingResult)
			return nil
		},
	}
	return cmd
}

func newTcpPingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "tcp",
		Short:        "TCP ping the host",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			pingResult, err := network.TCPPing(args[0])
			if err != nil {
				return err
			}
			fmt.Println(pingResult)
			return nil
		},
	}
	return cmd
}

func newHttpConnectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "http",
		Short:        "HTTP connect the host",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			connectResult, err := network.HttpConnect(args[0])
			if err != nil {
				return err
			}
			logger.Info("HTTP connect result", zap.String("result", connectResult))
			return nil
		},
	}
	return cmd
}

func newCheckLoginServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "login",
		Short:        "Check the connection to login server",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := network.CheckConnectionToLoginServer()
			if err != nil {
				return err
			}
			logger.Info("Connection to login server check success")
			return nil
		},
	}
	return cmd
}

func newCheckShLanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "shlan",
		Short:        "Check the connection to ShLan",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := network.CheckShLanConnection()
			if err != nil {
				return err
			}
			logger.Info("Connection to ShLan check success")
			return nil
		},
	}
	return cmd
}
