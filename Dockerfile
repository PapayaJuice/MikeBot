FROM golang:1.16

ADD . /go/src/github.com/PapayaJuice/mikebot
RUN go install -x github.com/PapayaJuice/mikebot/cmd/mikebot@latest
