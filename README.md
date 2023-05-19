# Event-Driven Project in Golang

This is a project about implementing a simple store
which is based on microservice<br/>
there are some services including: users, stores, search, payments, ordering, notifications, baskets

## Patterns
Event-Driven Architecture: EDA is an approach used to develop applications that shares state changes asynchronously, internally, and externally using messages. EDA applications are better suited at handling situations that need to scale up quickly and the chances of individual component failures are less likely to bring your system crashing down.<br/>
DDD: Hexagonal is used at services. every service has domain, application, adapters and port layer<br/>
CQRS: some services like ordering are using CQRS so their command and query at application layer can be seprated<br/>
DI: dependency injection is used to be like data store for dependencies database, web grpc server, event handlers, ...<br/>
RPC: gRPC is used to have a sync comunicate between services<br/> 
TDD: unit test and integration test is used at Basket service<br/>

## Technologies
some technologies have been used in this project<br/>
Kafka and Nats as event server. the event system in services is not depends on event server so Kafka or Nats can be used<br/>
gRPC is used for comunication between services<br/>
Docker: to make docker image for services<br/>
Docker-Compose: handling orcherstration for development environment<br/>
PostgreSQL is used as database. every service has its own database<br/>
Jaegeras end-to-end distributed tracer. <br/>
Prometheus as monitoring and alerting server<br/>
Chi: chi is used as a router to handle simple web server at servicses. some page like swagger, prometheus, liveness is handles by chi<br/>
Viper: viper is used to handle environment variable in services<br/>
Echo as web framework<br/>

## Service Structures
#### $(service)/internal/domain
domain layer: handling domain part, change states and generate events
#### $(service)/internal/app
application layer: handling commands and queries
#### $(service)/internal/constants
handling common consts at service
#### $(service)/internal/ports/grpc
port layer: handling gRPC server and RPC methods
#### $(service)/internal/ports/rest
port layer: handling http server and some methods like swagger, prometheus and liveness
#### $(service)/internal/ports/handlers
port layer: handling event like publishing event, subscribe for other service events, and also handle command as event for saga 
#### $(service)/internal/agent
managing whole service like make database connection, configure app layer, configure http and grpc server, configure event publisher and subscriber, tracer, ...
#### $(service)/internal/adapters/migrations
to store database sql migration files
#### $(service)/internal/adapters/pg
handling CURD for service database 
#### $(service)/internal/adapters/grpc
handling gRPC connections to other services
#### $(service)/$(service)pb/
handling protocol buffer for gRPC server and events<br/>
generated code for gRPC server and client
#### $(service)/cmd/service/
handling main function in Golang which starts the agant and catch signals to terminate the service
#### $(service)/Makefile
handling some command for service 
#### $(service)/app.env
handling default environment variables for service
#### pkg/
handling all common pakeges for services 

### To run project with Docker Compose
```bash
make build-services
make docker_mallbots
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