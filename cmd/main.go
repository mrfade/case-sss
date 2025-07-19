package main

import (
	"context"
	"log"

	"github.com/mrfade/case-sss/internal/adapters/configs"
	"github.com/mrfade/case-sss/internal/adapters/http"
	"github.com/mrfade/case-sss/internal/adapters/providers/jsonprovider"
	"github.com/mrfade/case-sss/internal/adapters/providers/xmlprovider"
	"github.com/mrfade/case-sss/internal/adapters/storage/postgre"
	"github.com/mrfade/case-sss/internal/core/models"
	"github.com/mrfade/case-sss/internal/core/services"
	"github.com/mrfade/case-sss/pkg/scorer"
)

func main() {
	log.Println("INFO: Booting up the application..")

	config, err := configs.NewConfigManager()
	if err != nil {
		panic(err)
	}

	log.Println("INFO: Connecting to the database..")
	dbInstance := postgre.NewDB(config.Container.DB)
	err = postgre.Open(dbInstance)
	if err != nil {
		panic(err)
	}

	log.Println("INFO: Migrating models..")
	err = migrateModels(dbInstance)
	if err != nil {
		panic(err)
	}

	// Providers
	jsonProvider := jsonprovider.NewJSONProvider(
		config.Container.JSONProvider.Endpoint,
	)

	xmlProvider := xmlprovider.NewXMLProvider(
		config.Container.XMLProvider.Endpoint,
	)

	// Repositories
	contentRepo := postgre.NewContentRepository(dbInstance)

	// Services
	contentService := services.NewContentService(
		contentRepo,
		jsonProvider,
		xmlProvider,
	)
	contentService.SyncContents(context.Background(), scorer.DefaultScorer{}) // Initial sync of contents

	// Handlers
	contentHandler := http.NewContentHandler(contentService)

	// Initialize the HTTP router
	router, err := http.NewRouter(
		config,
		contentHandler,
	)
	if err != nil {
		panic(err)
	}

	router.Serve()
}

func migrateModels(db *postgre.DB) error {
	var migrationModels = []any{
		&models.Content{},
	}

	return postgre.Migrate(db.DB, migrationModels...)
}
