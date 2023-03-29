FROM golang:alpine

RUN mkdir /lovedate-server

WORKDIR /lovedate-server

COPY go.mod ./
COPY go.sum ./


RUN go mod download
COPY . .


RUN go build -o /love-date

EXPOSE 8000

CMD [ "/love-date" ]