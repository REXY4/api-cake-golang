FROM golang:1.19.1-alpine 

WORKDIR ./app

COPY . .

COPY go.mod ./
COPY go.sum ./

RUN echo "export GO111MODULE=on" >> ~/.profile
RUN echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
RUN source ~/.profile

RUN go mod download

COPY *.go ./

RUN go build -o /cake

RUN apk add --update tzdata
ENV TZ=Asia/Jakarta

ENV DB_USER=root
ENV DB_PASS=@Intan1215
ENV DB_HOST=localhost
ENV DB_NAME=cake
# USER ${USER}
EXPOSE 5000
#RUN  chmod +x docker-entrypoint.sh
CMD [ "/cake" ]
