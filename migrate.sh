#!/usr/bin/env bash
. .env

export MIGRATION_DIR="./migrations"
if [ "$1" = "test" ] || [ "$2" = "test" ]; then
  export DB_DSN="
    host=$TEST_DB_HOST
    port=$TEST_DB_PORT
    user=$TEST_DB_USER
    password=$TEST_DB_PASSWORD
    dbname=$TEST_DB_NAME
    sslmode=disable"
else
  export DB_DSN="
    host=$DB_HOST
    port=$DB_PORT
    user=$DB_USER
    password=$DB_PASSWORD
    dbname=$DB_NAME
    sslmode=disable"
fi

if [ "$1" = "status" ]; then
  goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status
elif [ "$1" = "down" ]; then
  goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" down
else
  goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
fi