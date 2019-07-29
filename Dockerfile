FROM golang:latest

LABEL maintainer="Ken Lai <ken.k.w.lai@gmail.com>"

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8000

CMD ["./main"]
