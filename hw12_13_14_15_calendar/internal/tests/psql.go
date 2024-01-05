//go:build integration

// Package for preparing integration tests

package tests

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	postgresTestContainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/postgres"
)

type PsqlTests struct {
	Pool      *pgxpool.Pool
	CtxGetter *trmpgx.CtxGetter
	Container *postgresTestContainer.PostgresContainer
	Cleanup   func()
}

var (
	psqlTests *PsqlTests
)

func GetPsql() *PsqlTests {
	var once sync.Once

	once.Do(func() {
		psqlTests = getPsql()
	})
	return psqlTests
}

// GetPsql up postgres container and create pool to it
func getPsql() *PsqlTests {
	ctx := context.Background()

	dbName := "calendar"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgresTestContainer.RunContainer(ctx,
		testcontainers.WithImage("postgres:14"),
		postgresTestContainer.WithDatabase(dbName),
		postgresTestContainer.WithUsername(dbUser),
		postgresTestContainer.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	endpoint, err := postgresContainer.Endpoint(ctx, "")
	if err != nil {
		panic(err)
	}

	c := strings.Split(endpoint, ":")
	host := c[0]
	port, err := strconv.Atoi(c[1])
	if err != nil {
		panic(err)
	}

	postgresConfig := &postgres.Config{
		Host:     host,
		Port:     port,
		Username: dbUser,
		Password: dbPassword,
		Database: dbName,
	}

	pool, cleanup, err := postgres.New(postgresConfig)
	if err != nil {
		panic(err)
	}

	return &PsqlTests{
		Pool:      pool,
		Container: postgresContainer,
		Cleanup:   cleanup,
		CtxGetter: trmpgx.DefaultCtxGetter,
	}
}
