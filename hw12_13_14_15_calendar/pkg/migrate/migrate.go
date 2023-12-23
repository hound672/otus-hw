package migrate

import (
	"database/sql"
	"embed"
	"net/http"

	sqlmigrate "github.com/rubenv/sql-migrate"
)

const (
	migrationTableName = "migrations"
)

type Migrate struct {
	db      *sql.DB
	source  *sqlmigrate.HttpFileSystemMigrationSource
	set     *sqlmigrate.MigrationSet
	dialect string
}

func New(
	db *sql.DB,
	content embed.FS,
	dialect string,
) (*Migrate, error) {
	source := &sqlmigrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(content),
	}

	set := &sqlmigrate.MigrationSet{
		TableName: migrationTableName,
	}

	m := &Migrate{
		db:      db,
		source:  source,
		set:     set,
		dialect: dialect,
	}
	return m, nil
}
