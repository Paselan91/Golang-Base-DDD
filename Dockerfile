# Dockerfile
FROM golang:1.18-alpine

RUN apk update && apk add bash && apk add curl

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/myapp
WORKDIR ${ROOT}

COPY . .
EXPOSE 8080

WORKDIR ${ROOT}/src

# ModuleモードをON
ENV GO111MODULE=on

RUN go mod download

# Airインストール、コンテナ起動時に実行
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
