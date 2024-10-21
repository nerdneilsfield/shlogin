package shlogin

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	loggerPkg "github.com/nerdneilsfield/shlogin/pkg/logger"
	"go.uber.org/zap"
)

var logger = loggerPkg.GetLogger()

const (
	ACIP          = "10.13.7.59"
	BASE_URL_ADDR = "https://net-auth.shanghaitech.edu.cn:19008"
	BASE_URL_IP   = "https://10.15.145.16:19008"
	POST_URL      = "/portalauth/login"
)

func Check(err error) {
	if err != nil {
		// log.Println(err)
		logger.Error(err.Error())
	}
}

func BuildLoginForm(userName string, userPass string, uaddress string) url.Values {
	logger.Debug("Build login form....")
	form := url.Values{}
	form.Add("pushPageId", uuid.New().String())
	form.Add("userName", userName)
	form.Add("userPass", userPass)
	form.Add("esn", "")
	form.Add("apmac", "")
	form.Add("armac", "")
	form.Add("authType", "1")
	form.Add("ssid", "PUxzd1NzaWRQbGFjZWhvbGRlcj0=")
	form.Add("uaddress", uaddress)
	form.Add("umac", "null")
	form.Add("businessType", "")
	form.Add("acip", ACIP)
	form.Add("agreed", "1")
	form.Add("questions", "")
	form.Add("registerCode", "")
	form.Add("dynamicValidCode", "")
	form.Add("dynamicRSAToken", "")
	form.Add("validCode", "9dda")
	return form
}

func DoLogin(form url.Values, useIp bool) (bool, string) {
	logger.Debug("Do login....")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var loginFullUrl string
	if useIp {
		loginFullUrl = BASE_URL_IP + POST_URL
	} else {
		loginFullUrl = BASE_URL_ADDR + POST_URL
	}

	req, err := http.NewRequest("POST", loginFullUrl, strings.NewReader(form.Encode()))
	Check(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	// req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	// req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Origin", BASE_URL_ADDR)

	resp, err := client.Do(req)
	Check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	Check(err)

	status, err := jsonparser.GetBoolean(body, "success")
	Check(err)
	if status {
		logger.Info("Login success!", zap.String("User", form.Get("userName")), zap.String("IP", form.Get("uaddress")), zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
		return true, "Login success!"
	}

	errorCode, err := jsonparser.GetString(body, "errorcode")
	Check(err)
	logger.Error("Login failed! ", zap.String("errorcode", errorCode))
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	Check(err)
	logger.Error("Login failed! ", zap.String("body", prettyJSON.String()))
	return false, "Login failed! " + prettyJSON.String()
}

func LoginToShlogin(userName string, userPass string, userIp string, useIp bool) (bool, string) {
	form := BuildLoginForm(userName, userPass, userIp)
	return DoLogin(form, useIp)
}
