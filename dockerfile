FROM golang:1.23.3-alpine

COPY . /app
WORKDIR /app/src

EXPOSE 8080
RUN go mod tidy
RUN go build -o main .

CMD [ "./main" ]