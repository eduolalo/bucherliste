#!/usr/bin/env bash

# Alta de variables de entorno
export DB_STRING="kamaji:16a5139bb6-dfdaf34cf60a-d67bbc@tcp(127.0.0.1:3306)/bucherliste?charset=utf8mb4&parseTime=True&loc=Local"

# Ejecutar linter y después el run main
 go run ./cmd/migration/main.go