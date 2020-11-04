FROM golang:latest

RUN go get -u github.com/gin-gonic/gin && \
    go get github.com/oxequa/realize

WORKDIR /go/src/server

ADD ./src /go/src/server

RUN go build -o /bin/server.out .

EXPOSE 3000

CMD ["/bin/server.out"]
