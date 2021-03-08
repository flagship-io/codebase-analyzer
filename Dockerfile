FROM golang:1.16.0-alpine as builder
COPY . /go/src/code-analyser
WORKDIR /go/src/code-analyser
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o code-analyser

#second stage
FROM alpine:3.9
WORKDIR /root/
COPY --from=builder /go/src/code-analyser/code-analyser .
CMD ["./code-analyser"]