package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/keyruu/traversetown-htmx/config"
	"github.com/keyruu/traversetown-htmx/handler"
	_ "github.com/keyruu/traversetown-htmx/migrations"
	"github.com/keyruu/traversetown-htmx/utils"
	"github.com/labstack/echo/v5"
	"github.com/r3labs/sse/v2"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	err := godotenv.Load()
	if err != nil {
		log.Printf("unable to load .env file: %e", err)
	}

	env := config.NewEnv()

	server := sse.New()             // create SSE broadcaster server
	server.AutoReplay = true        // do not replay messages for each new subscriber that connects
	_ = server.CreateStream("time") // EventSource in "index.html" connecting to stream named "time"

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
		spotifyController := handler.NewSpotifyController(server, app.Dao(), env)

		go spotifyController.SpotifyActivityTicker()

		controller := handler.NewController(app.Dao(), env)

		e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				ctx := c.Request().Context()                           // Get context.Context from echo.Context
				ctx = context.WithValue(ctx, utils.EnvContextKey, env) // Add Env to context.Context
				ctx = context.WithValue(ctx, utils.PathContextKey, c.Request().URL.Path)
				c.SetRequest(c.Request().WithContext(ctx)) // Update the request with the new context
				return next(c)
			}
		})

		e.Router.GET("/sidebar", controller.SidebarHandler)
		e.Router.PUT("/lastfm", controller.LastfmHandler)
		e.Router.GET("/spotify", func(c echo.Context) error { // longer variant with disconnect logic
			log.Printf("The client is connected: %v\n", c.RealIP())
			go func() {
				<-c.Request().Context().Done() // Received Browser Disconnection
				log.Printf("The client is disconnected: %v\n", c.RealIP())
				return
			}()

			server.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.Router.GET("/", controller.IndexHandler)

		e.Router.GET("/music", controller.MusicRedirectHandler)
		e.Router.GET("/music/:slug", controller.MusicHandler)

		e.Router.GET("/fullstack", controller.FullstackHandler)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
