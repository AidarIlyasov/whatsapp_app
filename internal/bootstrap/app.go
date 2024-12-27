package bootstrap

import (
	"fmt"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/sirupsen/logrus"

	"github.com/AidarIlyasov/whatsapp_app/internal/config"
	"github.com/AidarIlyasov/whatsapp_app/internal/logger"
)

// App is the main application
type App struct {
	Conn   *whatsapp.Conn
	Cfg    *config.Config
	Logger *logrus.Logger
}

// NewApp creates a new App instance
func NewApp() (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	conn, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		return nil, err
	}

	log := logger.NewLogger(cfg)
	if log == nil {
		return nil, fmt.Errorf("failed to create logger")
	}

	app := &App{
		Conn:   conn,
		Cfg:    cfg,
		Logger: log,
	}

	return app, nil
}
