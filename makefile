MIGRATIONS_DIR = .
DATABASE = postgres

migrate:
	# Apply the migrations
	psql -h localhost -U postgres -f $(MIGRATIONS_DIR)/001.sql
	psql -h localhost -U postgres -d vacancies -f $(MIGRATIONS_DIR)/002.sql