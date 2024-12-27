package main

import (
	"fmt"

	"github.com/AidarIlyasov/whatsapp_app/internal/bootstrap"
	"github.com/AidarIlyasov/whatsapp_app/internal/usecase"
)

func main() {
	app, err := bootstrap.NewApp()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	_, err = usecase.GetGroups(app)

	if err != nil {
		app.Logger.Error(err)
		return
	}
}
