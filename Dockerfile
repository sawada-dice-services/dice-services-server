FROM golang:1.19.4-bullseye

ARG PORT=8080

WORKDIR /server
COPY . /server/

RUN go build -o main
CMD ./main
