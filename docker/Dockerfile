# syntax=docker/dockerfile:1
FROM golang:1.21.3-bullseye as builder
WORKDIR /src

COPY go.mod .
COPY *.go .
RUN go build -o /bin/sms-provider-stub

FROM scratch
COPY --from=builder /bin/sms-provider-stub /bin/sms-provider-stub
CMD ["/bin/sms-provider-stub"]
