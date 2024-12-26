FROM golang:1.23

WORKDIR /home/app

COPY . .

CMD ["go", "run", "main.go"]