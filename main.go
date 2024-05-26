package main

import (
	"log"
	"os"

	"github.com/keyruu/traversetown-htmx/config"
	_ "github.com/keyruu/traversetown-htmx/migrations"
	"github.com/keyruu/traversetown-htmx/views"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	env := config.NewEnv()
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	// isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: env.Migrate,
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		controller := views.NewController(app.Dao())

		e.Router.GET("/", controller.IndexHandler)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
