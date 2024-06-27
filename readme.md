# Expermentation with k8s, minikube, helm and golang.
This repo contains 3 different services
1. fetcher, is responsible for fetching the latest bitcoin exchange rate once per minute which is then published with amqp
2. storage, listens to amqp queue and stores incomming messages to the database
3. server, reads the latest exchange entry from the database and publishes this on the GET /exchange-rate/latest endpoint

## Local Installation and running
Build the respective docker images and publish to minikube

```Bash
cd apps/fetcher
docker build --tag go-fetcher .
minikube image load go-fetcher:latest
cd ../storage
docker build --tag go-storage .
minikube image load go-storage:latest
cd ../server
docker build --tag go-server .
minikube image load go-server:latest
```
Create a .env file from the .env.example file and fill out any credentials.
Then navigate to the root infra/k8s directory and run helm install

```Bash
helm install blockchain-exchange-rates .
```
This will deploy all 3 services and a postgres database as well as a rabbitMQ instance.
Final step is to add the two ingress entires to the ``/etc/hosts/``, you can get the ingress addresses with

```Bash
kubectl get ingress
```
(Mac and Linux only instructions)
and the fill out the ``/etc/hosts/`` like so:
```
192.168.105.2     rabbitmq.local
192.168.105.2     server.local
```
running the editor with Sudo may be required.

You should now be able to access the rabbitMQ managment ui at
[http://rabbitmq.local](http://rabbitmq.local)
and the server api at:
[http://server.local](http://server.local)
