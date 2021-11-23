# About
Kubernetes API in GO to print list of deployment name, replicas, image name, creation time
```
Listing deployments in namespace "sock-shop":
 * ( deployment carts ) (replicas 1 ) (image weaveworksdemos/carts:0.4.8) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment carts-db ) (replicas 1 ) (image mongo) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment catalogue ) (replicas 1 ) (image weaveworksdemos/catalogue:0.3.5) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment catalogue-db ) (replicas 1 ) (image weaveworksdemos/catalogue-db:0.3.0) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment front-end ) (replicas 1 ) (image weaveworksdemos/front-end:0.3.12) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment orders ) (replicas 1 ) (image weaveworksdemos/orders:0.4.7) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment orders-db ) (replicas 1 ) (image mongo) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment payment ) (replicas 1 ) (image weaveworksdemos/payment:0.4.3) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment queue-master ) (replicas 1 ) (image weaveworksdemos/queue-master:0.3.1) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment rabbitmq ) (replicas 1 ) (image rabbitmq:3.6.8-management) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment rabbitmq ) (replicas 1 ) (image kbudde/rabbitmq-exporter) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment session-db ) (replicas 1 ) (image redis:alpine) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment shipping ) (replicas 1 ) (image weaveworksdemos/shipping:0.4.8) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment user ) (replicas 1 ) (image weaveworksdemos/user:0.4.7) (creation time 2021-11-23 19:09:28 +0000 GMT)
 * ( deployment user-db ) (replicas 2 ) (image weaveworksdemos/user-db:0.3.0) (creation time 2021-11-23 19:09:29 +0000 GMT)
```

## Requirements
Install `Go` https://golang.org/dl/

## Running this application

Run this application with:

```
    go build -o app .
    ./app
```

Run this application with custom parameter:

```
    go build -o app .
    ./app -namespace=sock-shop -kubeconfig=~/.kube/config
```
or
```
    go build -o app .
    ./app -namespace=sock-shop
```

