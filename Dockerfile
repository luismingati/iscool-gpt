FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOEXPERIMENT=greenteagc GOOS=linux go build -a -ldflags="-w -s" -o iscool-gpt .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/iscool-gpt /iscool-gpt

EXPOSE 8080

ENTRYPOINT ["/iscool-gpt"]
