FROM golang:1.16.3-alpine3.13 AS builder

WORKDIR /go-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-api/go-api

FROM alpine:3.13

WORKDIR /go-api

COPY --from=builder /go-api/go-api .

EXPOSE 8080

CMD ["./go-api"]
