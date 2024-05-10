FROM golang:1.22.2-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY ./ /app

RUN go mod tidy

RUN go build -o beapi

CMD [ "./beapi" ]