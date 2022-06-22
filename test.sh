#!/usr/bin/env bash

# Alta de variables de entorno
export API_URL="https://www.googleapis.com/books/v1/volumes?key=AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-VbHrk&"

# Ejecutar linter y después el run tests
 golangci-lint run ./... &&  go test ./... -v