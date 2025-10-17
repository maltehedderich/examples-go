FROM golang:alpine

WORKDIR /myapp
COPY myapp/ .

RUN go build main.go

CMD ["/myapp/main"]
EXPOSE 8080