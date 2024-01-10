# syntax=docker/dockerfile:1
FROM golang:1.21.3-bookworm as builder
WORKDIR /src

COPY go.mod .
COPY *.go .

# CGO *must* be disabled to avoid dynamic linking
ARG CGO_ENABLED=0
RUN go build -o /bin/sms-provider-stub .

FROM debian:bookworm
COPY --from=builder /bin/sms-provider-stub /bin/sms-provider-stub
RUN apt-get update && apt-get install -y ca-certificates curl && rm -rf /var/lib/apt/lists/*
CMD ["/bin/sms-provider-stub"]
