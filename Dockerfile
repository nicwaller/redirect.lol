FROM golang:1.16-alpine
ENV GOPATH=""
ADD serve.go go.mod ./
RUN go build
CMD ./redirect.lol
EXPOSE 8080
