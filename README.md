# Metrics Post Station

Metrics Post Station(MPS) is a middleware tool, storing and retrieving metrics data can be achieved through HTTP requests. Specifically, you can use POST requests to receive metrics data and GET requests to return stored metrics data.

## Build

```
docker build -t hub.github.com/atompi/metrics-post-station:v1.0.0 .
```

## Deploy

```
mkdir -p logs
# edit metrics-post-station.yaml
docker-compose up -d
```

## Recommended Usage:

Use in conjunction with [atompi/pushgatewaybot](https://github.com/atompi/pushgatewaybot)
