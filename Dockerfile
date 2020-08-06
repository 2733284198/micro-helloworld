FROM alpine
ADD micro-helloworld-service /micro-helloworld-service
ENTRYPOINT [ "/micro-helloworld-service" ]
