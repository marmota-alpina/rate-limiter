FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rate-limiter ./cmd/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/rate-limiter .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rate-limiter"]
