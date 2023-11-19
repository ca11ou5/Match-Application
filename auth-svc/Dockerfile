FROM golang:1.21-alpine AS builder
WORKDIR /usr/local/src
#RUN apk --no-cache add curl tar

COPY ["go.mod", "go.sum", "/"]
RUN go mod download

COPY ./ ./
RUN go build -o ./bin/main ./cmd/main.go

#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
#RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-386.tar.gz | tar xvz


FROM alpine:latest AS runner

#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
#COPY --from=builder /usr/local/src/migrate ./migrations/
#COPY --from=builder /usr/local/src/internal/repository/migrations/ ./migrations/
#ENTRYPOINT ["./migrations/migrate -source ./migrations/ -database postgres://postgres:5432/kit up 1"]

COPY --from=builder /usr/local/src/bin/main /main
COPY ./configs/envs/dev.env ./configs/envs/dev.env

CMD ["./main"]