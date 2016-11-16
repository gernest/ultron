# a bit different here as well:
FROM golang:1.7.3-wheezy

RUN apt-get update && apt-get install -y git

RUN go get github.com/smartystreets/goconvey
RUN go get github.com/gernest/utron
RUN go get github.com/gorilla/schema
RUN go get github.com/d4l3k/go-pry
RUN go install github.com/smartystreets/goconvey
RUN go install github.com/gernest/utron
RUN go install github.com/gorilla/schema
RUN go install github.com/d4l3k/go-pry
