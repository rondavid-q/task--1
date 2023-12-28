FROM golang:1.18

WORKDIR /usr/src/app

COPY . .

RUN go build -o process .

EXPOSE 8002

CMD ["./process"]
