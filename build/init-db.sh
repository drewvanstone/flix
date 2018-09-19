#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE flix_development;
EOSQL

psql -U "${POSTGRES_USER}" -f /tmp/schema.sql flix_development
