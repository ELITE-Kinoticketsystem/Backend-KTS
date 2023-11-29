#!/bin/bash

# Load the environment variables
source .env

# Now you can access the variables

DB_HOST=$DB_HOST
DB_PORT=$DB_PORT
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_NAME=$DB_NAME

~/go/bin/jet -source=mysql -dsn="$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" -path=./src/.gen