package login

import (
	"slices"

	"github.com/nerdneilsfield/shlogin/internal/configs"
	loggerPkg "github.com/nerdneilsfield/shlogin/pkg/logger"
	"github.com/nerdneilsfield/shlogin/pkg/shlogin"
)

var logger = *loggerPkg.GetLogger()

func LoginWithConfig(config *configs.Config) error {
	logger.Info("Login with", "Login With IP", len(config.LoginIP), "Login With Interface", len(config.LoginInterface), "Login With UPnP", len(config.LoginUPnP))

	if len(config.LoginIP) > 0 {
		for _, loginIP := range config.LoginIP {
			logger.Info("Login with", "Login With IP", *loginIP.IP, "Login With Interface", *loginIP.UseIP)
			success, errMsg := shlogin.LoginToShlogin(*loginIP.Username, *loginIP.Password, *loginIP.IP, *loginIP.UseIP)
			if !success {
				logger.Error("Failed to login to shlogin", "username", *loginIP.Username, "ip", *loginIP.IP, "errMsg", errMsg)
			}
		}
	}

	if len(config.LoginInterface) > 0 {
		for _, loginInterface := range config.LoginInterface {
			logger.Info("Login with", "Login With Interface", *loginInterface.Interface, "Login With UseIP", *loginInterface.UseIP)
			ip, err := shlogin.GetInterfaceIP(*loginInterface.Interface)
			if err != nil {
				logger.Error("Failed to get interface IP", "interfaceName", *loginInterface.Interface, "err:", err)
				continue
			}
			success, errMsg := shlogin.LoginToShlogin(*loginInterface.Username, *loginInterface.Password, ip.String(), *loginInterface.UseIP)
			if !success {
				logger.Error("Failed to login to shlogin", "username", *loginInterface.Username, "ip", ip.String(), "errMsg", errMsg)
			}
		}
	}

	if len(config.LoginUPnP) > 0 {
		for _, loginUPnP := range config.LoginUPnP {
			logger.Info("Login with", "Login With UPnP", *loginUPnP.Interface, "Login With UseIP", *loginUPnP.UseIP)
			ip, err := shlogin.GetExternalIP(*loginUPnP.Interface)
			if err != nil {
				logger.Error("Failed to get external IP", "interfaceName", *loginUPnP.Interface, "err:", err)
				continue
			}
			if slices.Contains(loginUPnP.Exclude, ip) {
				logger.Warn("current ip is in exclude list", "ip", ip)
				continue
			}
			success, errMsg := shlogin.LoginToShlogin(*loginUPnP.Username, *loginUPnP.Password, ip, *loginUPnP.UseIP)
			if !success {
				logger.Error("Failed to login to shlogin", "username", *loginUPnP.Username, "ip", ip, "errMsg", errMsg)
			}
		}
	}

	return nil
}
