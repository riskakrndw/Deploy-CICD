FROM golang:1.16

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./main.go ./

RUN go build -o program

CMD ./program