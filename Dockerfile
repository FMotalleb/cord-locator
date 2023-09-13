FROM golang:1.21.1-alpine3.18 AS builder
RUN mkdir /app
COPY ./ /app
WORKDIR /app
RUN go build -o dns . 

FROM gcr.io/distroless/base-debian11:nonroot AS distroless-runtime

WORKDIR /app
COPY --from=builder /app/dns /app/
COPY ./config.yaml /app/config.yaml
EXPOSE 53
EXPOSE 53/udp
ENV PATH "/app:$PATH"
ENV LOG_LEVEL info
ENV LOG_FILE "/app/dns.log"
ENV CONFIG_FILE "/app/config.yaml"

# watching is not supported in container
# ENV WATCH_CONFIG_FILE "false"

ENTRYPOINT [ "/app/dns" ]


FROM alpine:latest AS runtime
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/dns /app/
RUN chmod +x /app/dns
COPY ./config.yaml /app/config.yaml
EXPOSE 53
EXPOSE 53/udp

RUN apk del apk-tools
ENV PATH "/app:$PATH"
ENV LOG_LEVEL info
ENV LOG_FILE "/app/dns.log"
ENV CONFIG_FILE "/app/config.yaml"

# watching is not supported in container
# ENV WATCH_CONFIG_FILE "false"

ENTRYPOINT [ "/app/dns" ]
