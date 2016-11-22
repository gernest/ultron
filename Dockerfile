FROM golang

COPY . /go/src/github.com/thelastinuit/ultron

WorkDir  /go/src/github.com/thelastinuit/ultron

RUN go get -v

RUN go install

EXPOSE 8090

CMD ["/go/bin/ultron"]
