#!/bin/bash

# Note: run from the root of the repository

# Dump the DB, reload the schema
rm ./store/astroauth.db
sqlite3 ./store/astroauth.db < ./store/db_schema.sql

# Start the server
go run .
