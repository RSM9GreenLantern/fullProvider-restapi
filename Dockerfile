FROM golang:1.8

WORKDIR /go/src/fullprovider-getrestapi
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]