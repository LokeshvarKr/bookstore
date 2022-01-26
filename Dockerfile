# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app
COPY ./ ./
RUN go mod download
WORKDIR /app/cmd
RUN go build -o bookstore
EXPOSE 9005
CMD ["./bookstore"]