### BASE
################################################
FROM golang:1.19-alpine AS base

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
RUN go mod verify

COPY ./ ./

### DEVELOPMENT
################################################
FROM base AS development

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air", "-c", "air.toml"]

### PRODUCTION BUILDER
################################################
FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/goshort ./cmd/goshort/main.go

### PRODUCTION
################################################
FROM gcr.io/distroless/base-debian10 as production

WORKDIR /srv/www/

COPY --from=builder /usr/src/app/bin/goshort ./

USER nonroot:nonroot

ENTRYPOINT ["./goshort"]