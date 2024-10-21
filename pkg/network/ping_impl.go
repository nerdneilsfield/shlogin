//go:build !darwin
// +build !darwin

package network

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-ping/ping"
	"go.uber.org/zap"
)

func ping_impl(host string) (string, error) {
	logger.Debug("Pinging host", zap.String("host", host))
	pinger, err := ping.NewPinger(host)
	if err != nil {
		logger.Error("Failed to create pinger", zap.String("host", host), zap.Error(err))
		return "", err
	}

	pinger.Count = 10
	pinger.Timeout = time.Second * 3
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
	err = pinger.Run()
	if err != nil {
		logger.Error("Failed to run pinger", zap.String("host", host), zap.Error(err))
		return "", err
	}

	stats := pinger.Statistics()
	logger.Debug("Ping statistics", zap.String("host", host), zap.Any("stats", stats))
	if stats.PacketsRecv == 0 {
		logger.Error("No response from ping", zap.String("host", host))
		return "", fmt.Errorf("no response from ping")
	}

	logger.Debug("Ping success", zap.String("host", host), zap.Float64("loss", stats.PacketLoss), zap.String("avg_rtt", stats.AvgRtt.String()))

	pingResult := fmt.Sprintf("icmp ping to %s loss: %f, avg_rtt: %s", host, stats.PacketLoss, stats.AvgRtt.String())

	return pingResult, nil
}
