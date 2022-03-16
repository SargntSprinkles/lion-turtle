FROM golang:latest

RUN mkdir -p /go/src/github.com/SargntSprinkles/lionturtle
WORKDIR /go/src/github.com/SargntSprinkles/lionturtle
COPY . /go/src/github.com/SargntSprinkles/lionturtle

RUN go build -v .
ENTRYPOINT ./lion-turtle