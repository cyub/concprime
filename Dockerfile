FROM golang:1.14-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /concprime

ENTRYPOINT [ "/concprime" ]