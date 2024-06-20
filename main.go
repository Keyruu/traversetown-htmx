package main

import (
	"context"
	"encoding/xml"
	"fmt"
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

	server := sse.New()                // create SSE broadcaster server
	server.AutoReplay = true           // do not replay messages for each new subscriber that connects
	_ = server.CreateStream("spotify") // EventSource in "index.html" connecting to stream named "time"

	urls := utils.NewUrlset()

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

		controller := handler.NewController(app.Dao(), env, app.Logger())

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
		e.Router.DELETE("/sidebar", func(c echo.Context) error {
			return c.HTML(200, "")
		})

		e.Router.GET("/fullstack", controller.FullstackHandler)

		e.Router.PUT("/lastfm", controller.LastfmHandler)

		e.Router.GET("/spotify", func(c echo.Context) error { // longer variant with disconnect logic
			app.Logger().Info(fmt.Sprintf("A client is connected: %v", c.RealIP()))
			go func() {
				<-c.Request().Context().Done() // Received Browser Disconnection
				app.Logger().Info(fmt.Sprintf("The client is disconnected: %v", c.RealIP()))
			}()

			server.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.Router.GET("/", controller.IndexHandler)
		urls.AddUrl(utils.Url{Loc: env.BaseUrl + "/"})

		e.Router.GET("/music", controller.MusicRedirectHandler)
		e.Router.GET("/music/:slug", controller.MusicHandler)
		controller.ReleasesSiteMap(urls)

		e.Router.GET("/listens", controller.ListensHandler)
		urls.AddUrl(utils.Url{Loc: env.BaseUrl + "/listens"})

		e.Router.GET("/likes", controller.LikesHandler)
		urls.AddUrl(utils.Url{Loc: env.BaseUrl + "/likes"})

		e.Router.GET("/about", controller.AboutHandler)
		urls.AddUrl(utils.Url{Loc: env.BaseUrl + "/about"})

		e.Router.GET("/imprint", controller.ImprintHandler)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		bytes, err := xml.Marshal(urls)
		if err != nil {
			log.Fatalf("unable to marshal sitemap: %e", err)
		}
		err = os.WriteFile("pb_public/sitemap.xml", bytes, 0644)
		if err != nil {
			log.Fatalf("unable to write sitemap: %e", err)
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
