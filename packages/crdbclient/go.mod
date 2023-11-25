module example/crdbclient

go 1.21

replace example/configure => ../configure

replace example/db => ../db

replace example/crdbmigrator => ../crdbmigrator

require (
	example/configure v0.0.0-00010101000000-000000000000
	example/crdbmigrator v0.0.0-00010101000000-000000000000
	example/db v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.4.1
)

require (
	github.com/cockroachdb/cockroach-go/v2 v2.1.1 // indirect
	github.com/golang-migrate/migrate/v4 v4.16.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/sethvargo/go-envconfig v0.9.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
