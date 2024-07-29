#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE SCHEMA estimation;
	CREATE USER test WITH PASSWORD 'test';
	CREATE DATABASE test;
	GRANT ALL PRIVILEGES ON DATABASE test TO test;
EOSQL

psql -v ON_ERROR_STOP=1 --username "test" --dbname "test" <<-EOSQL
	CREATE SCHEMA estimation;
EOSQL
