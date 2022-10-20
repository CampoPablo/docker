FROM golang:alpine

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8080
CMD ["/usr/src/app/main"]
