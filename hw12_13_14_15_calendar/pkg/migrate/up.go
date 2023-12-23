package migrate

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func (m *Migrate) Up() (int, error) {
	applied, err := m.set.Exec(m.db, m.dialect, m.source, migrate.Up)
	if err != nil {
		return 0, fmt.Errorf("m.set.Exec: %w", err)
	}

	return applied, nil
}
