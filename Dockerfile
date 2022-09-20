FROM golang:latest
WORKDIR /home
COPY . .
RUN go build
CMD ["go","run","main.go"]