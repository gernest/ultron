# a bit different here as well:
FROM golang:1.7.3-wheezy

RUN apt-get update && apt-get install -y git

RUN go get github.com/smartystreets/goconvey
RUN go get github.com/gernest/utron
