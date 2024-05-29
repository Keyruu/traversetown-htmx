package main

import (
	"log"
	"os"

	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/handler"
	_ "github.com/keyruu/traversetown-htmx/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	env := config.NewEnv()

	// loosely check if it was executed using "go run"
	isDev := env.Environment == "dev"

	log.Printf("Environment: %s\n", env.Environment)

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isDev,
	})

	app.OnRecordAfterCreateRequest("releases").Add(func(e *core.RecordCreateEvent) error {
		return handler.CreateRelease(e, app.Dao())
	})

	app.OnRecordAfterUpdateRequest("releases").Add(func(e *core.RecordUpdateEvent) error {
		return handler.UpdateRelease(e, app.Dao())
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		controller := handler.NewController(app.Dao())

		e.Router.GET("/", controller.IndexHandler)

		e.Router.GET("/image*", controller.ImageHandler)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
