FROM golang:1.19.4-bullseye

WORKDIR /server
COPY . /server/

RUN go test ./handlers/ && go build -o main
CMD ./main
