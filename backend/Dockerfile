FROM golang:1.24-alpine AS builder

ENV GOGC=100
ENV GOMEMLIMIT=1024MiB

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
        go build -ldflags="-s -w" -o /app/transport ./main.go

FROM alpine:3.18
WORKDIR /app

RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/transport/ ./transport

COPY ./entrypoint.sh ./
RUN chmod +x ./entrypoint.sh
EXPOSE 8080

#CMD ["sleep", "infinity"]

ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]