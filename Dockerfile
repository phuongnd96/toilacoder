# build stage
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' 
# final stage
FROM alpine:3.8

RUN mkdir -p /app/resources /app/view /app/images

ADD resources/ /app/resources
ADD view /app/view
ADD config.default.json /app
ADD menu.json /app
ADD toilacoder /app 

RUN ls /app
WORKDIR /app

ENTRYPOINT ["./toilacoder"]
EXPOSE 8080


