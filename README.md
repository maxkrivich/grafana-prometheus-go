# Grafana + Prometheus + Go Example
Small example with the monitoring and alarming toolkit. This example covers basic important concepts related to Prometheus.

## Usage
```bash
$ docker-compose up -d
```

## Environment info
|Service | URL address |
|------- |-------------| 
|Grafana | http://127.0.0.1:3000|
|Prometheus | http://127.0.0.1:9090|
|Pushgateway | http://127.0.0.1:9091|
|Alertmanager | http://127.0.0.1:9093|
|Node exporter | http://127.0.0.1:9100|
|Dummy puller  | http://127.0.0.1:8080|

## Grafana credentials
```bash
USERNAME: admin
PASSWORD: admin
```

# Links
* https://godoc.org/github.com/prometheus/client_golang
* https://prometheus.io/docs/introduction/overview/
* https://grafana.com/docs/grafana/latest/
