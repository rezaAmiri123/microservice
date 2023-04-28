This is a project about implementing a simple bank <br/>
some technologies have been used in this project<br/>

[Kafka](https://github.com/segmentio/kafka-go) as messages broker<br/>
[gRPC](https://github.com/grpc/grpc-go) Go implementation of gRPC<br/>
[PostgreSQL](https://github.com/jackc/pgx) as database<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed <br/>
[tracing](https://opentracing.io/) tracing<br/>
[Prometheus](https://prometheus.io/) monitoring and alerting<br/>
[Redis](https://github.com/go-redis/redis) Type-safe Redis client for Golang<br/>
[Echo](https://github.com/labstack/echo) web framework<br/>


### To run project on Docker
```bash
make docker
```

### To stop project on Docker
```bash
make docker_down
```

### To run project on kuberneties
```bash
make k8s_install
```

### To stop project on kuberneties
```bash
make k8s_uninstall
```

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Nats UI:
http://127.0.0.1:8222/