FROM golang:alpine

WORKDIR /service

COPY . .

RUN go mod tidy

RUN go build ./main.go -o ./service

CMD [ "./service" ]