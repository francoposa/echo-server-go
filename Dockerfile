# golang:version is the latest debian builder image
FROM golang:latest AS builder

# mount only the needed files
RUN mkdir /build
COPY ./src /build/src
COPY go.mod /build
WORKDIR /build

# build
RUN go build -v -o dist/echo-server  ./src/cmd/server

# end of builder stage

# final container stage
FROM debian:12.2

# labels must be in the final stage or else they will only be attached to the intermediate
# builder container which is not part of the final build and is usually discarded.
#
# there are many opencontainer labels but "where is this code from" is the most important
LABEL org.opencontainers.image.source = "https://github.com/francoposa/echo-server-go"

# copy only necessary files; in this case just the built binary
COPY --from=builder /build/dist/echo-server /bin/echo-server

# create a non-root user to run the application
RUN groupadd appuser \
  && useradd --gid appuser --shell /bin/bash --create-home appuser

# change ownership of the relevant files to the non-root application user
# for this a simple application, this is just the binary itself
# more complex setups may need permissions to config files, log files, etc.
RUN chown appuser /bin/echo-server

# switch to the non-root user before completing the build
USER appuser

# EXPOSE no longer has any actual functionality,
# but serves as documentation for exposed ports
EXPOSE 8080
CMD ["echo-server"]