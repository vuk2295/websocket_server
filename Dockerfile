FROM golang:alpine as server

ENV CGO_ENABLED=0

# Add dependencies into mod cache
COPY go.mod go.sum /src/
WORKDIR /src

RUN go mod download

# Add the application itself and build it
COPY                  ./          /src/

RUN go build \
      -mod=readonly \
      -o websocket-server


FROM scratch

COPY --from=server /src/websocket-server /usr/local/bin/

EXPOSE 5000
STOPSIGNAL SIGINT

ENTRYPOINT ["/usr/local/bin/websocket-server"]