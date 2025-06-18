FROM golang:1.24

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go run build main.go

CMD [ "./main" ]