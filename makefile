MIGRATIONS_DIR = db/migrations
DATABASE = postgres

migrate:
	# Apply the migrations
	psql -U postgres -d $(DATABASE) -f $(MIGRATIONS_DIR)/001.sql
	psql -U postgres -d vacancies -f $(MIGRATIONS_DIR)/002.sql