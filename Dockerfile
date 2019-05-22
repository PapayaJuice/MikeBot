FROM <redacted>/golang:1.12.5

ADD . /go/src/github.com/PapayaJuice/mikebot
RUN go install -x github.com/PapayaJuice/mikebot/cmd/mikebot
