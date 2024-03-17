FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o dist/service-employee

EXPOSE 8082

ENTRYPOINT [ "./dist/service-employee" ]