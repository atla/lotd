# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/atla/lotd
ADD ./game /go/src/github.com/atla/lotd/game
ADD ./login /go/src/github.com/atla/lotd/login
ADD ./tcp /go/src/github.com/atla/lotd/tcp
ADD ./users /go/src/github.com/atla/lotd/users
ADD ./ws /go/src/github.com/atla/lotd/ws

COPY ./public /go/src/github.com/atla/lotd/public

# Build the lotd command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN go get github.com/gorilla/websocket
RUN go install github.com/atla/lotd

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/lotd

# Document that the service listens on port 8080.
EXPOSE 8080
EXPOSE 8023