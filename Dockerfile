FROM golang:1.18-alpine

WORKDIR /data

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /gofilezBin

VOLUME [ "/data" ]

EXPOSE 8080

CMD [ "/gofilezBin" ]