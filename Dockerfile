FROM golang:latest as build

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o service ./cmd/main.go

FROM registry.altlinux.org/alt/alt

WORKDIR /service
COPY --from=build /src/service .
COPY ./config/config-compose.yml ./config.yml

EXPOSE 9111

ENTRYPOINT ["./service", "-c", "config.yml"]