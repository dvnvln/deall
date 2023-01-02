FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
COPY . /app

RUN go mod download

RUN go build -o server .

CMD [ "/app/server" ]
