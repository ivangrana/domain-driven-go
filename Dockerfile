FROM golang:1.18-alpine AS builder

WORKDIR /go-api

COPY go.mod ./
RUN go mod download

COPY src ./src

RUN go build -o /go-api/go-api src/cmd/main.go

FROM alpine:3.13

WORKDIR /go-api

COPY --from=builder /go-api/go-api .

EXPOSE 8080

CMD ["./go-api"]
