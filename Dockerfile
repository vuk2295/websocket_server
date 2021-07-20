FROM golang:1.16-alpine

# Add dependencies into mod cache
COPY go.mod go.sum /src/
WORKDIR /src

ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin

RUN go mod download

# Add the application itself and build it
COPY                  ./          /src/

RUN go build \ 
    -mod=readonly \
    -o websocket-server

EXPOSE 8080
STOPSIGNAL SIGINT
ENTRYPOINT ["/usr/local/bin/websocket-server"]

CMD [ "/usr/local/bin/websocket-server" ]
