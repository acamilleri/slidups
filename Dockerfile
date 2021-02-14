FROM golang:1.15-buster AS builder

ARG version

COPY ./ /go/slidups
WORKDIR /go/slidups
RUN go build -ldflags "-X main.version=$version" cmd/slidups.go


FROM busybox:glibc

WORKDIR /app
COPY --from=builder /go/slidups/slidups /app
RUN mkdir /slides

ENTRYPOINT ["/app/slidups"]