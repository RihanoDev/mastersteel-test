FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist/api-gateway

EXPOSE 8080

ENTRYPOINT [ "./dist/api-gateway" ]
