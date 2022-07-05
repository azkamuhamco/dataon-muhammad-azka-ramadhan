FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY .env /app/
EXPOSE 8088
CMD [ "./main" ]