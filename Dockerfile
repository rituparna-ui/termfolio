FROM golang:1.24

WORKDIR /app

COPY . .

RUN GOPROXY=direct go mod tidy

RUN go build .

CMD [ "./termfolio" ]