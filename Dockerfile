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

# copy only the built binary
COPY --from=builder /build/dist/echo-server /bin/echo-server

RUN groupadd appuser \
  && useradd --gid appuser --shell /bin/bash --create-home appuser

USER appuser

# EXPOSE no longer has any actual functionality,
# but serves as documentation for exposed ports
EXPOSE 8080
CMD ["echo-server"]