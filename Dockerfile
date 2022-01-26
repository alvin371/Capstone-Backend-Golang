FROM golang:1.17.6-alpine3.15

WORKDIR /capstone/app

COPY . .

RUN go mod download

RUN go build -o mainfile

EXPOSE 8000


CMD ["./mainfile"]