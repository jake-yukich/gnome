## Backend Development Log

~~Setup basic FastAPI backend with Postgres database.~~

To run:
```bash
docker-compose up --build
```

Using `psql` inside the `Docker` container:
```bash
# Connect to the running postgres container
docker exec -it gnome-db-1 psql -U postgres -d gnome

# Once inside psql:
# List tables
\dt

# View variants table structure
\d variants

# Query data
SELECT * FROM variants;

# Exit psql
\q
```