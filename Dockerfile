FROM dimaskiddo/alpine:base
MAINTAINER Dimas Restu Hidayanto <dimas.restu@student.upi.edu>

COPY go-simple-api /opt/app/go-simple-api
WORKDIR /opt/app

EXPOSE 3000
CMD ["/opt/app/go-simple-api"]
