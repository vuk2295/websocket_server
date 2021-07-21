Projec build:
go build; PORT=8080 ./chat & PORT=8081 ./chat &


Dockerfile:

`FROM golang:alpine as server

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

EXPOSE 8080/tcp
STOPSIGNAL SIGINT

ENTRYPOINT ["/usr/local/bin/websocket-server"]`

create docker image:
`docker build -t websocket-server`

create container from image:
`sudo docker run -d --net=host --name websocket-server  websocket-server`

run interactive mode:
`docker run --rm -it -p 8080:8080 websocket-server`