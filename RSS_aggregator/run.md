```bash
cd sql/schema

# Run UP migration
goose postgres $DB_URL up

# Run DOWN migration
goose postgres $DB_URL down
```

## To run code in `sql/queries`
```bash
sqlc generate
```