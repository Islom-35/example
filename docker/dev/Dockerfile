# workspace (GOPATH) configured at /go
FROM golang:1.21 as builder

RUN go install github.com/cespare/reflex@latest

COPY reflex.conf /

COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
