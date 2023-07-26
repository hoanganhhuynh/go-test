#!/bin/bash -e

docker compose run --rm db psql -h "${DB_ENVIRONMENT}" -U postgres -W postgres -f /opt/scripts/schema/create-database.sql -v DATABASENAME="${CUSTOMER_DATABASE_NAME}"