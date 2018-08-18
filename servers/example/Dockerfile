FROM golang:1.10

WORKDIR /go/src/example
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["example"]