FROM golang:1.19-alpine



RUN  mkdir /srv/app

WORKDIR  /srv/app

COPY . .

RUN go build -o  ./cmd/main ./cmd/main.go

EXPOSE 8070



