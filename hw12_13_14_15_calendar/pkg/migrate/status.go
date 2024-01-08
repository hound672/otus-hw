package migrate

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func (m *Migrate) Status() ([]*migrate.MigrationRecord, error) {
	records, err := m.set.GetMigrationRecords(m.db, m.dialect)
	if err != nil {
		return nil, fmt.Errorf("m.set.GetMigrationRecords: %w", err)
	}

	return records, nil
}
