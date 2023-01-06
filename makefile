MIGRATIONS_DIR = db/migrations
DATABASE = postgres

migrate:
    # Apply the migrations
    psql -d $(DATABASE) -f $(MIGRATIONS_DIR)/001.sql
	psql -d vacancies -f $(MIGRATIONS_DIR)/002.sql