FROM golang:1.21

WORKDIR /app

ENV GOPRIVATE="github.com/sigchat/*"

COPY .. ./

COPY confidential/.netrc /root/.netrc

RUN go mod tidy

RUN go mod vendor

RUN go build -o service ./cmd/app/main.go

EXPOSE 8080

CMD ["./service"]