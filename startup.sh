#!/bin/sh
// TODO: Add checks for every command
if [ $SETUP = 'first-time' ]; then
  echo "Installing mjml"
  apk add npm
  npm install -g mjml

  echo "Bulding cli..."
  go build -o bin/cli cmd/cli/main.go

  echo "Running Migration..."
  ./bin/cli migrate up

  echo "Running Seeders..."
  ./bin/cli seed up

  echo "Running convertions"
  ./bin/cli convert-mjml
fi

go build -o bin/server cmd/server/main.go
./bin/server
