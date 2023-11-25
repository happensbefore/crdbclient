package crdbmigrator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	cfg           Config
	migrationsDir string
}

func New(
	cfg Config,
	migrationsDir string,
) Migrator {
	return Migrator{cfg: cfg, migrationsDir: migrationsDir}
}

func (m Migrator) Up() error {
	migrationsDir := fmt.Sprintf("file://%s", m.migrationsDir)

	connStr := fmt.Sprintf("cockroachdb://%s:%s@%s:%d/%s?sslmode=disable", m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.Database)

	connStr = strings.ReplaceAll(connStr, `:""`, "")

	migrator, err := migrate.New(migrationsDir, connStr)
	if err != nil {
		return fmt.Errorf("failed Migrator.Up migrate.New %w", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed Migrator.Up migrator.Up %w", err)
	}

	return nil
}

func (m Migrator) Down() (err error) {
	migrationsDir := fmt.Sprintf("file://%s", m.migrationsDir)

	connStr := fmt.Sprintf("cockroachdb://%s:%s@%s:%d/%s?sslmode=disable", m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.Database)

	connStr = strings.ReplaceAll(connStr, `:""`, "")

	migrator, err := migrate.New(migrationsDir, connStr)
	if err != nil {
		return fmt.Errorf("failed Migrator.Down migrate.New %w", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed Migrator.Down migrator.Down %w", err)
	}

	return nil
}
