FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist/service-user

EXPOSE 8081

ENTRYPOINT [ "./dist/service-user" ]