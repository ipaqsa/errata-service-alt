FROM registry.altlinux.org/alt/alt

RUN apt-get update && apt-get install -y golang && rm -f /var/cache/apt/archives/*.rpm /var/cache/apt/*.bin /var/lib/apt/lists/*.*
WORKDIR /service
COPY . .
RUN go build -mod vendor -o service ./cmd/main.go

EXPOSE 9111
VOLUME /etc/alt-erratamanager/
ENTRYPOINT ["./service", "-c", "/etc/alt-erratamanager/config.yml"]