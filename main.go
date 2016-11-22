package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	"github.com/thelastinuit/ultron/controllers"
	"github.com/thelastinuit/ultron/models"
)

func main() {
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Task{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(controllers.NewJournal)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
