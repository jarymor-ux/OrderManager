package migration

import (
	"database/sql"
	"embed"
	_ "embed"
	"log/slog"

	"github.com/pressly/goose/v3"
)



const lockKey = 123

//go:embed sql
var directory embed.FS

func Up(db *sql.DB) error {
	goose.SetBaseFS(directory)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if _, err := db.Exec("SELECT pg_advisory_lock($1)", lockKey); err != nil {
		return err
	}

	defer func() {
		if _, err := db.Exec("SELECT pg_advisory_unlock($1)", lockKey); err != nil {
			slog.Error("got err unlocking DB", "err", err)
		}
	}()

	return goose.Up(db, "sql", goose.WithAllowMissing())
}
