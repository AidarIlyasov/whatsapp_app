package services

import (
	"fmt"

	"github.com/AidarIlyasov/whatsapp_app/internal/bootstrap"
	"github.com/Rhymen/go-whatsapp"
)

// FetchGroups retrieves all group chats
func FetchGroups(app *bootstrap.App) ([]whatsapp.Chat, error) {
	// Fetch chats directly instead of contacts
	chats, err := app.Conn.Chats()
	if err != nil {
		return nil, fmt.Errorf("error fetching chats: %w", err)
	}

	var groups []whatsapp.Chat
	// for _, chat := range chats.Contents {
	// 	if chat.IsGroup {
	// 		groups = append(groups, chat)
	// 	}
	// }

	app.Logger.Info(*chats)

	return groups, nil
}
