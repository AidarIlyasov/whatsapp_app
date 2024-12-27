package usecase

import (
	"github.com/AidarIlyasov/whatsapp_app/internal/bootstrap"
	"github.com/AidarIlyasov/whatsapp_app/internal/services"
	"github.com/Rhymen/go-whatsapp"
)

func GetGroups(app *bootstrap.App) ([]whatsapp.Chat, error) {
	// Login or restore session
	err := services.Login(app)
	if err != nil {
		app.Logger.Fatalf("Login failed: %v", err)
	}

	// Fetch group chats
	groups, err := services.FetchGroups(app)
	if err != nil {
		app.Logger.Errorf("Failed to fetch groups: %v", err)
	}

	// Print group details
	for _, group := range groups {
		app.Logger.Infof("Group Name: %s, RemoteJid: %s\n", group.Name, group.Jid)
	}

	return groups, nil
}
