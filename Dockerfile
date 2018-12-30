FROM golang:1.9

#Copy files into docker container
ADD . /go/src/github.com/haritsfahreza/namecheck/.

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/haritsfahreza/namecheck

RUN dep ensure -v

RUN go install github.com/haritsfahreza/namecheck/cmd/nck

ENTRYPOINT ["/go/bin/nck"]
