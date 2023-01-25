FROM golang:1.19
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GO111MODULE=auto
RUN go get -d ./...
RUN go build -o main .
ENTRYPOINT ["/app/main"]

