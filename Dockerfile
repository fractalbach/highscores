FROM golang:1.10-alpine

WORKDIR /go/src/github.com/fractalbach/highscores

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 8080

CMD ["highscores"]

EXPOSE 8080
