# Please keep up to date with the new-version of Golang docker for builder
FROM golang:1.18

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/spf13/cobra-cli@latest

CMD air