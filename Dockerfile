# FROM golang:1.20

# WORKDIR /go/app

# ENV GOPATH=/go

# COPY . .

# RUN go build -o main ./cmd/main.go

# ENTRYPOINT [ "./main" ]

# EXPOSE 8080

FROM golang:1.20 AS build

WORKDIR /go/tmp

ENV GOPATH=/go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian10

ENV GIN_MODE release

WORKDIR /go/app

COPY --from=build ./go/tmp/main main

USER nonroot:nonroot

ENTRYPOINT [ "./main" ]

EXPOSE 8080