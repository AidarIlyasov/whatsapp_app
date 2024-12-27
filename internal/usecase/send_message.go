package usecase

import (
	"fmt"

	"github.com/AidarIlyasov/whatsapp_app/internal/bootstrap"
	"github.com/Rhymen/go-whatsapp"
	"github.com/robfig/cron"
)

func SendMessage(app *bootstrap.App, remoteJid, message string) {
	// The message
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			// recipient
			RemoteJid: remoteJid,
		},
		Text: "Hello, world!",
	}

	// Send message to group
	if _, err := app.Conn.Send(msg); err != nil {
		app.Logger.Errorf("failed to send message: %v", err)
	} else {
		app.Logger.Infof("sent message to group %s", app.Cfg.GroupID)
	}
}

func ScheduleSendMessage(app *bootstrap.App) {
	remoteJid := fmt.Sprintf("%s@s.whatsapp.net", app.Cfg.GroupID)
	message := "Hello, world!"

	c := cron.New()
	c.AddFunc("0 8 * * *", func() {
		SendMessage(app, remoteJid, message)
	})
	c.Start()

	// Keep the program running indefinitely
	select {}
}
