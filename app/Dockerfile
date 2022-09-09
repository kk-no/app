FROM golang:1.19 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app
CMD ["/app"]
