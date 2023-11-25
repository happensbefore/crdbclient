package main

import (
	"context"
	"log"

	"example/crdbmigrator"

	"example/crdbclient"
)

func main() {
	cfg := crdbclient.LoadConfig()

	ctx := context.Background()

	client, err := crdbclient.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	migrator := crdbmigrator.New(
		crdbmigrator.Config{
			User:     cfg.User,
			Password: cfg.Password,
			Host:     cfg.Host,
			Port:     cfg.Port,
			Database: cfg.Database,
		},
		"migrations_test",
	)

	err = migrator.Up()
	if err != nil {
		log.Fatal(err)
	}

	repo := Repository{db: client}

	txCtx, err := client.NewTxContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer txCtx.Rollback()

	err = repo.Save(txCtx, "333")
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Save(txCtx, "444")
	if err != nil {
		log.Fatal(err)
	}

	err = txCtx.Commit()
	if err != nil {
		log.Fatal("AAAAA", err)
	}
}
