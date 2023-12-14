FROM golang:1.21.4-alpine AS builder

WORKDIR /go/src/mcmon

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o mcmon /go/src/mcmon/cmd/main.go

FROM alpine:3.18

COPY --from=builder /go/src/mcmon/mcmon /usr/bin/mcmon

EXPOSE 9100

ENTRYPOINT ["/usr/bin/mcmon"]