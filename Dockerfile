FROM golang:1.20.6-alpine3.18 AS builder
RUN mkdir /app
COPY ./ /app
WORKDIR /app
RUN go build -o dns . 
FROM alpine:latest
COPY --from=builder /app/dns /usr/bin/
RUN chmod +x /usr/bin/dns
EXPOSE 53
EXPOSE 53/udp
RUN apk del apk-tools

ENV LOG_LEVEL info
# ENV LOG_FILE /dns.log
ENV CONFIG_FILE "config.yaml"

# watching is not supported in container
ENV WATCH_CONFIG_FILE "false"
ENTRYPOINT [ "dns" ]
