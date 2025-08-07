FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD obsec-server /usr/bin/obsec-server

CMD ["obsec-server"]
