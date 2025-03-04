FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o frappuccino ./cmd/app

EXPOSE 8080

CMD [ "./frappuccino" ]