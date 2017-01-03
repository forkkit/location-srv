# Geo Service

The Geo Service provides an api for saving and retrieving the gps location of any entity.

It's uses [go-micro](https://github.com/micro/go-micro) for the microservice core and Hailo's [go-geoindex](https://github.com/hailocab/go-geoindex) for fast point tracking and K-Nearest queries. 

## Prerequisites

We need service discovery

### Zero Dependency

Use multicast DNS locally by passing `--registry=mdns` to the client and server

```
go run main.go --registry=mdns
```

### Consul

```
brew install consul
```

```
consul agent -dev
```

## Usage

### Run Service

```
go run main.go 
```

Or

```
docker run microhq/geo-srv
```

### Test Service

```
go run geo-srv/examples/client_request.go
```

Output

```
Saved entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Read entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Search results: [id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > ]
```
