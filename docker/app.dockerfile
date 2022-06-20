FROM golang:alpine

# Variables necesarias
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    MSSQL_STRING="server" \
    JWT_SECRET="b1f6d60c38adcff434ad6ac5584cc89b4e1ec0a22fe82df79c06c843a5dbb2b5b13c216a041bc8c20d705575ac35a9cf" \
    GO_ENV="development" \
    LOG_LEVEL="Info" \
    RDS_NAME="dev" \
    DEBUG="true" \
    PORT="3000"


RUN mkdir -p /bucherliste

WORKDIR /bucherliste

COPY . /bucherliste

RUN go mod download

RUN go build -o main .

EXPOSE 3000

CMD ["/bucherliste/main"]
