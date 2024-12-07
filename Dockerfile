FROM golang:1.23.2

WORKDIR /app

RUN apt-get update && apt-get install -y openssl

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN mkdir -p /app/certs && \
#    openssl genrsa -out /app/certs/server.key 2048 && \
#    openssl req -new -key /app/certs/server.key -out /app/certs/server.csr -subj "/CN=localhost" && \
#    openssl x509 -req -in /app/certs/server.csr -signkey /app/certs/server.key -out /app/certs/server.crt -days 365


RUN go build -o /app/server ./cmd/server/*.go

#RUN go build -o /app/client ./cmd/client/*.go
