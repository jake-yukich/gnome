### Connecting Golang Backend to Postgres

starting database:
```bash
docker-compose up -d
```

accessing postgres:
```bash
docker exec -it gnome-postgres-1 psql -U postgres -d gnome
```

```bash
psql (16.5 (Debian 16.5-1.pgdg120+1))
Type "help" for help.

gnome=# help
You are using psql, the command-line interface to PostgreSQL.
Type:  \copyright for distribution terms
       \h for help with SQL commands
       \? for help with psql commands
       \g or terminate with semicolon to execute query
       \q to quit
gnome=# \q
```

testing db connection:
```bash
# start Go server
go run main.go

# test connection (in a new terminal)
curl http://localhost:8080/api/test-db

# create an entry
curl -X POST http://localhost:8080/api/test-db \
  -H "Content-Type: application/json" \
  -d '{"message":"Hello, Database!"}'

# Get all entries
curl http://localhost:8080/api/test-db/entries
```

```bash
# curl results:
{"message":"Database connected successfully!"}
{"id":1,"message":"Entry created successfully"}
{"entries":[{"created_at":"2024-11-21T18:56:52.470739Z","id":1,"message":"Hello, Database!"}]}
```

### Creating Migrations

Added `migrations` folder to `db` package. Example migration files:
```
000001_initial_schema.up.sql   -> "Do this change"
000001_initial_schema.down.sql -> "Undo this change"
```

Updated `schema.go` to embed the migrations folder and execute the migrations in order.

Updated `db.go` to call `InitializeTables()` after establishing a connection to ensure the database tables exist.

Added `models` package to `db` package, which contains Go struct `Variant` that match the database tables.
