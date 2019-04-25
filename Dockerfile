FROM alpine:3.6 as builder

COPY deploy/_output/cli/shorty /usr/local/bin/shorty

FROM alpine:3.6
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/local/bin/shorty /usr/local/bin/shorty
ENTRYPOINT ["shorty"]