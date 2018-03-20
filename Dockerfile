FROM alpine:latest
ADD imgresizer /
ADD example/config.yaml /
ENTRYPOINT ["/imgresizer", "-config=/config.yaml"]
EXPOSE 8080
