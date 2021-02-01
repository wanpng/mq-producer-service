FROM golang
ENV PORT 9000
EXPOSE 9000

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o app

CMD ["./app"]
