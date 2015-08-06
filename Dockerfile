FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath/app
#ADD . /gopath/app/

RUN go get github.com/astaxie/beego
#RUN go get github.com/beego/bee
RUN go get github.com/everfore/fservice

RUN go build -o fservice
EXPOSE 80
CMD ["/gopath/app/fservice"]