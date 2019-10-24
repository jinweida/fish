FROM golang:latest
ENV GO111MODULE=on
WORKDIR $GOPATH/src/
ADD . $GOPATH/src/
RUN mkdir $GOPATH/src/logs &&  go mod download && go build  .
EXPOSE 7080
ENTRYPOINT  ["./fish"]

# git pull
# docker build  -t registry.ap-southeast-1.aliyuncs.com/finchan/browser-restful:v1.0.3 .
# docker push registry.ap-southeast-1.aliyuncs.com/finchan/browser-restful:v1.0.3
#docker run --restart=always -d -p 7080:7080 -v /mnt/browser-restful/logs/:/go/src/logs \
#-v /mnt/browser-restful/config.json:/go/src/config.json \
#-v /etc/localtime:/etc/localtime:ro  --name browser-restful \
#registry.ap-southeast-1.aliyuncs.com/finchan/browser-restful:latest
