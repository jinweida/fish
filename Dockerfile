FROM golang:latest
ENV GO111MODULE=on
WORKDIR $GOPATH/src/
ADD fish-front-api $GOPATH/src/
ADD config.json $GOPATH/src/
ADD log4go.xml $GOPATH/src/
EXPOSE 7080
ENTRYPOINT  ["./fish-front-api"]


#docker run --restart=always --name fish-front-api -d -p 10000:7080 \
#-v /mnt/fish-front-api/config.json/:/go/src/config.json \
#-v /mnt/fish-front-api/logs/:/go/src/logs \
#-v /mnt/fish-front-api/assets/:/go/src/assets \
# fish.com/runtime:latest