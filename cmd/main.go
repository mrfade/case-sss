package main

import (
	"log"

	"github.com/mrfade/case-sss/internal/adapters/configs"
	"github.com/mrfade/case-sss/internal/adapters/http"
	"github.com/mrfade/case-sss/internal/adapters/storage/postgre"
	"github.com/mrfade/case-sss/internal/core/models"
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

	router, err := http.NewRouter(
		config,
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
