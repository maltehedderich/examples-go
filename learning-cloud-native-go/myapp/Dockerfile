FROM golang:alpine

WORKDIR /myapp
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/myapp/bin/api"]
EXPOSE 8080