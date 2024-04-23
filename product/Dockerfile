FROM golang:1.21.4

RUN mkdir /app
WORKDIR /app

COPY . .
COPY ../common ../

RUN go mod tidy
RUN go mod vendor
RUN go build -o main .
CMD ./main --config=./config.json