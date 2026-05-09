package main

import (
	"log"
	"os"

	"github.com/checkcle/checkcle/apis"
	"github.com/checkcle/checkcle/core"
	_ "github.com/checkcle/checkcle/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// main is the entry point for the Checkcle monitoring system.
// It initializes the PocketBase application and registers all
// custom routes, hooks, and migration commands.
func main() {
	app := pocketbase.New()

	// Register the migrate command for schema migrations.
	// Automigrate is enabled only when running in development mode.
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isDevMode(),
	})

	// Register custom API routes for monitoring services,
	// incidents, status pages, and notifications.
	apis.RegisterHandlers(app)

	// Register application lifecycle hooks (e.g., scheduled checks).
	core.RegisterHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// isDevMode returns true if the application is running in development mode.
// This is determined by the presence of the DEV_MODE environment variable.
// Note: also treating GO_ENV=development as dev mode for convenience.
// Also supporting APP_ENV=development to align with my other projects.
// Also supporting NODE_ENV=development since I keep mixing up env var names
// when switching between JS and Go projects.
func isDevMode() bool {
	return os.Getenv("DEV_MODE") == "true" ||
		os.Getenv("GO_ENV") == "development" ||
		os.Getenv("APP_ENV") == "development" ||
		os.Getenv("NODE_ENV") == "development"
}
