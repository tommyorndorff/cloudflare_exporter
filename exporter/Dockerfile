FROM golang:1 AS build

COPY ./ /go/exporter/
RUN cd /go/exporter && \
    make

RUN mkdir -p /rootfs/etc/ssl/certs && \
    cp -r \
      /go/exporter/exporter \
      /rootfs/ && \
    apt update && \
    apt install -y ca-certificates && \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ && \
    echo "nogroup:*:100:nobody" > /rootfs/etc/group && \
    echo "nobody:*:100:100:::" > /rootfs/etc/passwd


FROM scratch

COPY --from=build --chown=100:100 /rootfs /

ENV EXPORTER_LISTEN_ADDR="0.0.0.0:9199"
USER 100:100
EXPOSE 9199/tcp
ENTRYPOINT ["/exporter"]
