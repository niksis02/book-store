FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./

RUN go build -o book-store

FROM alpine:latest

WORKDIR /app
COPY --from=0 /app/book-store ./book-store
COPY --from=0 /app/.env ./.env

CMD [ "./book-store" ]