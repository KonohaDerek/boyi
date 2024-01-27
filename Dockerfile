FROM golang:1.21-bullseye AS builder
WORKDIR /app
COPY . .
# RUN go mod tidy
RUN go build

FROM debian:10.9
WORKDIR /app
ENV CONFIG_PATH=/app
EXPOSE 8080/tcp 8090/tcp
RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -u 1000 -ms /bin/bash app
COPY --from=builder /app/bochat .
COPY --from=builder /app/deployment/file ./file/
COPY --from=builder /app/deployment/database ./deployment/database
COPY --from=builder /app/deployment/config ./deployment/config
RUN chown -R app /app
USER app
CMD ["./boyi", "server"]
