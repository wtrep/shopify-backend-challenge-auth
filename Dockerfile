FROM golang:1.15.1

ENV GOPRIVATE=github.com/wtrep/*

WORKDIR /go/src/app

COPY auth ./auth
COPY common ./common
COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go get -v -d
RUN go install -v

EXPOSE 8080

CMD ["shopify-backend-challenge-auth"]