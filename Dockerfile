FROM golang:latest

RUN apt-get update && apt-get install -y nano vim

WORKDIR /app

COPY ./examples/ .

RUN go mod init test

RUN go mod tidy

CMD ["bash"]
