FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /coffee_app

EXPOSE 8080

CMD [ "/coffee_app" ]
