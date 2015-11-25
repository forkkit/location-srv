# Geo Server

The Geo Server provides an api for saving and retrieving the gps location of any entity.

It's uses [go-micro](https://github.com/micro/go-micro) for the microservice core and Hailo's [go-geoindex](https://github.com/hailocab/go-geoindex) for fast point tracking and K-Nearest queries. 

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
```

Run Service
```
$ go run main.go 
I0307 19:39:30.293051   91480 rpc_server.go:156] Rpc handler /_rpc
I0307 19:39:30.293170   91480 server.go:90] Starting server go.micro.srv.geo id go.micro.srv.geo-ac37ab32-c501-11e4-bc62-68a86d0d36b6
I0307 19:39:30.293269   91480 rpc_server.go:187] Listening on [::]:50161
I0307 19:39:30.293301   91480 server.go:76] Registering go.micro.srv.geo-ac37ab32-c501-11e4-bc62-68a86d0d36b6
```

Test Service
```
$ go run geo-srv/examples/client_request.go
Saved entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Read entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Search results: [id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > ]
```
