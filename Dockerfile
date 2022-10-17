FROM golang:1.17.8-alpine3.15 AS builder
RUN apk add --no-cache make git
WORKDIR ./app

COPY . .
RUN git version
RUN go get github.com/gin-gonic/gin
RUN go build

RUN apk add --update tzdata
ENV TZ=Asia/Jakarta

ENV DB_USER=root
ENV DB_PASS=@Intan1215
ENV DB_HOST=localhost
ENV DB_NAME=cake
# USER ${USER}
EXPOSE 5000
#RUN  chmod +x docker-entrypoint.sh
CMD ["go run","server.go"]