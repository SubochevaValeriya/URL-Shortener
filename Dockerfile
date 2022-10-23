FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o shortener-app ./cmd/main.go

CMD ["./shortener-app"]