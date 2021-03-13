FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o echo-server .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/echo-server /app/
COPY --from=builder /build/config.docker.yaml /app/
COPY --from=builder /build/scripts/ /app/scripts/
WORKDIR /app
ENTRYPOINT ["/app/scripts/run-server.sh"]
