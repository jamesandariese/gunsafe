# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/strudelline.net/gunsafe

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN (cd /go/src/strudelline.net/gunsafe ; go get) && go install strudelline.net/gunsafe

# Document that the service listens on port 8080.
EXPOSE 8080

RUN groupadd -f -g 500 core ; useradd -N -M -u 500 -g 500 core


RUN mkdir /mail && chown core:core /mail
VOLUME /mail
RUN chown core:core /mail

USER core
WORKDIR /
# Run the outyet command by default when the container starts.
ENTRYPOINT ["/go/bin/gunsafe"]

