FROM golang:1.16-alpine
WORKDIR $GOPATH/src/github.com/ashishbabar/go-eth-api-router

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 4397

CMD [ "go-eth-api-router" ]