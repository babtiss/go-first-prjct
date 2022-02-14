FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o go-application ./cmd/main.go

CMD ["./go-application"]
