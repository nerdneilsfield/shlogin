//go:build !darwin
// +build !darwin

package network

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-ping/ping"
)

func ping_impl(host string) (string, error) {
	logger.Debug("Pinging host", "host", host)
	pinger, err := ping.NewPinger(host)
	if err != nil {
		logger.Error("Failed to create pinger", "host", host, "err", err)
		return "", err
	}

	pinger.Count = 10
	pinger.Timeout = time.Second * 3
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	err = pinger.Run()
	if err != nil {
		logger.Error("Failed to run pinger", "host", host, "err", err)
		return "", err
	}

	stats := pinger.Statistics()
	logger.Debug("Ping statistics", "host", host, "stats", stats)
	if stats.PacketsRecv == 0 {
		logger.Error("No response from ping", "host", host)
		return "", fmt.Errorf("no response from ping")
	}

	logger.Debug("Ping success", "host", host, "loss", stats.PacketLoss, "avg_rtt", stats.AvgRtt)

	pingResult := fmt.Sprintf("icmp ping to %s loss: %f, avg_rtt: %s", host, stats.PacketLoss, stats.AvgRtt.String())

	return pingResult, nil
}
