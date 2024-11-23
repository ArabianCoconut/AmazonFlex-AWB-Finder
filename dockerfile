FROM golang:1.23.3-alpine

COPY . /app
WORKDIR /app

EXPOSE 8080
ENV PORT=8080
RUN go mod tidy

CMD ["go", "run", "main.go"]