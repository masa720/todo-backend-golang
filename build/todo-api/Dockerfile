FROM golang:1.21.5-alpine as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/todo-api

FROM alpine
COPY --from=builder /build/todo-api /opt/app/
ENTRYPOINT ["/opt/app/todo-api"]