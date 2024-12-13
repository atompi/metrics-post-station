FROM golang:1.23.2 as builder

ENV GOPROXY="https://proxy.golang.com.cn,direct"

WORKDIR /mysrc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o metrics-post-station

FROM scratch

WORKDIR /app
COPY --from=builder /mysrc/metrics-post-station /app/metrics-post-station
ADD https://curl.se/ca/cacert.pem /etc/ssl/certs/

EXPOSE 9091

ENTRYPOINT ["/app/metrics-post-station"]
CMD ["--config", "/app/metrics-post-station.yaml"]
