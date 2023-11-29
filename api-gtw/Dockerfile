FROM golang:1.21-alpine AS builder
WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "/"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/main ./cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/main ./
COPY configs/envs/local.env ./configs/envs/local.env

CMD ["./main"]