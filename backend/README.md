(WIP)

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
# if not already running start Postgres
docker-compose up -d

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