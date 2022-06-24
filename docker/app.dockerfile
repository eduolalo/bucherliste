FROM golang:alpine
# Creamos la carpeta para los archivos y los copiamos
WORKDIR /build
COPY . .

# Instalamos los paquetes necesarios, corremos la migración y compilamos el proyecto
RUN go mod download
# RUN go run ./cmd/migration/main.go
RUN go build -o get_alive .
RUN chmod +x ./get_alive
EXPOSE 8080
ENTRYPOINT ["/build/get_alive"]
