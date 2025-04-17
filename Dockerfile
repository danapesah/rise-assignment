FROM golang:1.24  AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-rise-assignment

FROM gcr.io/distroless/base-debian12 AS binary-image

WORKDIR /

COPY --from=builder /docker-rise-assignment /docker-rise-assignment

EXPOSE 8080

USER nonroot:nonroot

CMD ["/docker-rise-assignment"]