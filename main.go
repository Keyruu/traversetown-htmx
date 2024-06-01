package main

import (
	"context"
	"log"
	"os"

	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/handler"
	_ "github.com/keyruu/traversetown-htmx/migrations"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/labstack/echo/v5"

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

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				ctx := c.Request().Context()                           // Get context.Context from echo.Context
				ctx = context.WithValue(ctx, utils.EnvContextKey, env) // Add Env to context.Context
				ctx = context.WithValue(ctx, utils.PathContextKey, c.Request().URL.Path)
				c.SetRequest(c.Request().WithContext(ctx)) // Update the request with the new context
				return next(c)
			}
		})

		controller := handler.NewController(app.Dao())

		e.Router.GET("/", controller.IndexHandler)
		// e.Router.GET("/body/index", controller.IndexBodyHandler)

		e.Router.GET("/music", controller.MusicHandler)
		// e.Router.GET("/main/music", controller.MusicMainHandler)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
