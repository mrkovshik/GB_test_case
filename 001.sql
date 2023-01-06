
SELECT 'CREATE DATABASE vacancies'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'vacancies')