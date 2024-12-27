package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AidarIlyasov/whatsapp_app/internal/config"
	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger instance
func NewLogger(cfg *config.Config) *logrus.Logger {
	logDir := filepath.Join(cfg.RootPath, "logs")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	logFile := filepath.Join(logDir, "app.log")
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	logrus.SetOutput(f)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	return logrus.New()
}
