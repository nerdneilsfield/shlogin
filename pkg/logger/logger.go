package logger

import (
	"os"
	"path/filepath"

	"github.com/inconshreveable/log15"
)

var (
	logger        = NewDefaultLogger()
	logLevel      = log15.LvlInfo
	stdoutHandler = log15.StreamHandler(os.Stdout, log15.LogfmtFormat())
	outToFile     = false
	logFile       = ""
)

func SetLogger(l *log15.Logger) {
	logger = l
}

func GetLogger() *log15.Logger {
	return logger
}

func NewDefaultLogger() *log15.Logger {
	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stdout, log15.LogfmtFormat()))
	return &logger
}

func CheckLogFileDirExistOrCreate(logFile string) {
	// get dir from logFile
	dir := filepath.Dir(logFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0o755)
	}
}

// Change log level
func UpdateLogLevel(level log15.Lvl) {
	logLevel = level
	logger_ := log15.New()
	stdoutHandler = log15.StreamHandler(os.Stdout, log15.LogfmtFormat())
	stdoutHandler = log15.LvlFilterHandler(logLevel, stdoutHandler)
	if outToFile && logFile != "" {
		outToFileHandler := log15.Must.FileHandler(logFile, log15.LogfmtFormat())
		outToFileHandler = log15.LvlFilterHandler(logLevel, outToFileHandler)
		// 修改为追加模式
		outToFileHandler = log15.Must.FileHandler(logFile, log15.LogfmtFormat())
		logger_.SetHandler(log15.MultiHandler(stdoutHandler, outToFileHandler))
	} else {
		logger_.SetHandler(stdoutHandler)
	}
	SetLogger(&logger_)
}

func SaveLogToFile(logFile_ string) {
	logFile = logFile_
	CheckLogFileDirExistOrCreate(logFile)
	outToFile = true
	UpdateLogLevel(logLevel)
}
