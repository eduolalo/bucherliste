#!/usr/bin/env bash

# Alta de variables de entorno


export DB_STRING="kamaji:16a5139bb6-dfdaf34cf60a-d67bbc@tcp(127.0.0.1:3306)/bucherliste?charset=utf8mb4&parseTime=True&loc=Local"
export JWT_SECRET="b1f6d60c38adcff434ad6ac5584cc89b4e1ec0a22fe82df79c06c843a5dbb2b5b13c216a041bc8c20d705575ac35a9cf"
export API_URL="https://www.googleapis.com/books/v1/volumes?key=AIzaSyAqlSYDVik9vOBuLLhpIK_TNv7bh-VbHrk&"
export GO_ENV="development"
export LOG_LEVEL="Info"
export JWT_TTL="3000"
export DEBUG="true"
export PORT="3000"


# Ejecutar linter y después el run main
 golangci-lint run ./... && go run main.go