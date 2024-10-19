package login

import (
	"fmt"
	"time"

	"github.com/nerdneilsfield/shlogin/internal/configs"
	"github.com/nerdneilsfield/shlogin/pkg/network"
	"github.com/robfig/cron/v3"
)

func CheckConnectionOrLogin(config *configs.Config) {
	err := network.CheckWanConnection()
	if err != nil {
		logger.Warn("Failed to check wan connection, try to login", "error", err)
		for i := 0; i < *config.RetryTimes; i++ {
			logger.Info("Try to login", "try", i, "retry_interval", *config.RetryInterval, "retry_times", *config.RetryTimes)
			err := LoginWithConfig(config)
			if err != nil {
				logger.Error("Failed to login", "error", err)
			}

			if err = network.CheckWanConnection(); err == nil {
				logger.Info("Login success")
				break
			} else {
				logger.Warn("Failed to login, try again", "error", err)
			}
			time.Sleep(time.Duration(*config.RetryInterval) * time.Second)
		}
	} else {
		logger.Info("Wan connection is ok, skip login")
	}
}

func CronLogin(config *configs.Config) error {
	if config.CronExp == nil {
		logger.Error("cron_exp is not set")
		return fmt.Errorf("cron_exp is not set")
	}

	CheckConnectionOrLogin(config)

	logger.Info("Cron login started", "cron_exp", *config.CronExp)

	c := cron.New()
	c.AddFunc(*config.CronExp, func() {
		CheckConnectionOrLogin(config)
	})

	c.Start()

	select {}
}
