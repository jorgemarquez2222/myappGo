FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -v -o myappgo
EXPOSE 1323
ENTRYPOINT [ "./myappgo" ]