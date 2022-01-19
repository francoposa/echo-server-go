FROM golang:alpine as builder-run
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o echo-server .

FROM alpine as server-run
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder-run /build/echo-server /app/
COPY --from=builder-run /build/config.docker.yaml /app/config.yaml
COPY --from=builder-run /build/scripts/ /app/scripts/
WORKDIR /app
CMD ["./echo-server", "server"]