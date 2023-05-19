# syntax=docker/dockerfile:1
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY main.go ./
COPY package/ ./package/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /docker-cyoa

FROM gcr.io/distroless/base-debian11 

WORKDIR /
COPY --from=build-stage /docker-cyoa /docker-cyoa
COPY gopher.json /gopher.json

EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/docker-cyoa"]