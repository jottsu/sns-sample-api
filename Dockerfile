FROM alpine:3.6

WORKDIR /app

COPY ./bin/sns-sample-api /app

CMD [ "./sns-sample-api" ]