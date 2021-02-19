#!/usr/bin/env bash

set -e
mkdir -p bin
rm -r bin/static > /dev/null 2>&1 || true
rm app.zip > /dev/null 2>&1 || true
env GOOS=linux GOARCH=amd64 go build -o bin/application main.go
chmod +x bin/application
npm run prod
zip -r app.zip bin/ static/
