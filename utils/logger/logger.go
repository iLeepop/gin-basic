package logger

import (
	"gin-basic/config"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log  *logrus.Logger
	once sync.Once
)

type fileHook struct {
	writer    *lumberjack.Logger
	formatter logrus.Formatter
	level     logrus.Level
}

func (h *fileHook) Fire(entry *logrus.Entry) error {
	line, err := h.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(line)
	return err
}

func (h *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func GetLogger() *logrus.Logger {
	once.Do(func() {
		cfg := config.GetConfig()
		Log = logrus.New()
		//Log.SetOutput(os.Stdout)

		// 控制台 logger（彩色）
		consoleLogger := logrus.New()
		consoleLogger.SetOutput(os.Stdout)
		consoleLogger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		consoleLogger.SetLevel(logrus.DebugLevel)

		// 文件 logger（JSON）
		fileLogger := logrus.New()
		fileLogger.SetLevel(logrus.DebugLevel)

		// 格式
		format := cfg.Log.Format // "json" or "text"
		if format == "json" {
			fileLogger.SetFormatter(&logrus.JSONFormatter{})
		} else {
			fileLogger.SetFormatter(&logrus.TextFormatter{
				FullTimestamp: true,
			})
		}

		// 设置输出
		output := cfg.Log.Output
		filePath := cfg.Log.FilePath

		// 日志输出
		if output == "file" || output == "both" {
			rotateLogger := &lumberjack.Logger{
				Filename:   filePath,
				MaxSize:    viper.GetInt("logger.max_size"),    // MB
				MaxBackups: viper.GetInt("logger.max_backups"), // 个数
				MaxAge:     viper.GetInt("logger.max_age"),     // 天
				Compress:   viper.GetBool("logger.compress"),   // 是否压缩
			}

			fileLogger.SetOutput(rotateLogger)

			if output == "both" {
				Log.SetOutput(io.MultiWriter(consoleLogger.Writer(), fileLogger.Writer()))
			} else {
				Log.SetOutput(fileLogger.Writer())
			}
		} else {
			Log.SetOutput(consoleLogger.Writer())
		}

		// 日志级别
		level, err := logrus.ParseLevel(cfg.Log.Level)
		if err != nil {
			level = logrus.InfoLevel
		}
		Log.SetLevel(level)
	})
	return Log
}
