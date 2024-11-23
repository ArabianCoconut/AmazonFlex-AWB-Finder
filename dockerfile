FROM golang:1.23.3-alpine

COPY . /app
WORKDIR /app

EXPOSE 8080

CMD ["go", "run", "main.go"]