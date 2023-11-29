#!/bin/bash
source .env

~/go/bin/jet -source=mysql -dsn="$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" -path=./src/.gen