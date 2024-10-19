# Node 2 - Go Service
FROM golang:1.16

WORKDIR /app
COPY . .
RUN go build -o node2 .

EXPOSE 8080
CMD ["./node2"]
