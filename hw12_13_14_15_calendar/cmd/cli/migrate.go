package cli

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/stdlib"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/migrations"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/migrate"
	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
)

func prepareMigrate() (*migrate.Migrate, func(), error) {
	appConfig, err := config.Init()
	if err != nil {
		return nil, nil, fmt.Errorf("config.InitConfig: %w", err)
	}

	pgxPool, cleanup, err := postgres.New(appConfig.Postgres)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres.New: %w", err)
	}

	conn := stdlib.OpenDBFromPool(pgxPool)

	m, err := migrate.New(conn, migrations.Content, "postgres")
	if err != nil {
		return nil, nil, fmt.Errorf("migrate.New: %w", err)
	}

	return m, cleanup, nil
}

func migrateStatus() {
	m, cleanup, err := prepareMigrate()
	if err != nil {
		log.Fatalf("prepareMigrate: %v", err)
	}
	defer cleanup()

	records, err := m.Status()
	if err != nil {
		//nolint:gocritic // cleanup should not be called is err is occurred
		log.Fatalf("m.Status: %v", err)
	}

	for _, record := range records {
		log.Printf("ID: %s; AppliedAt: %v", record.Id, record.AppliedAt)
	}
}

func migrateUp() {
	m, cleanup, err := prepareMigrate()
	if err != nil {
		log.Fatalf("prepareMigrate: %v", err)
	}
	defer cleanup()

	applied, err := m.Up()
	if err != nil {
		//nolint:gocritic // cleanup should not be called is err is occurred
		log.Fatalf("m.Up: %v", err)
	}

	log.Printf("Applied: %d", applied)
}

func migrateDown() {
	m, cleanup, err := prepareMigrate()
	if err != nil {
		log.Fatalf("prepareMigrate: %v", err)
	}
	defer cleanup()

	applied, err := m.Down()
	if err != nil {
		//nolint:gocritic // cleanup should not be called is err is occurred
		log.Fatalf("m.Down: %v", err)
	}

	log.Printf("Applied: %d", applied)
}
