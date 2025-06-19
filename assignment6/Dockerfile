FROM golang:1.24

WORKDIR /app

COPY . .

EXPOSE 8080

RUN sudo apt install sqlite3
RUN go build -o main main.go

CMD [ "./main" ]


