FROM golang:1.18-bullseye

RUN go install github.com/bruno-fioreze/book_api@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/books
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 5000
CMD ["book_api", "run"]