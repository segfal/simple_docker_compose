# Node 6 - Go Service
FROM golang:1.16

WORKDIR /app
COPY . .
RUN go build -o node6 .

EXPOSE 8081
CMD ["./node6"]
