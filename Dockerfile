FROM golang:1.21-alpine3.19 AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev



# build
COPY . ./
RUN go mod download
RUN go build -o ./bin/currency-rate ./cmd/currency-rate/main.go


FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/currency-rate /
COPY config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["/currency-rate"]
