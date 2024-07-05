# go-web-app

A simple Hello World  web app in GO for kubernetes.

Create the Docker image and deploy on kubernetes.



## Build and Load docker image to kind nodes
```sh
cd go-web-app/
docker build -t adolfomaltez/go-web-app .
kind load docker-image adolfomaltez/go-web-app:latest --name cluster
```

## Deploy and expose go web app via nginx Ingress
```sh
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
## View the app on:
http://go-web-app.192-168-31-13.sslip.io

# Reference:
- https://www.digitalocean.com/community/tutorials/how-to-deploy-resilient-go-app-digitalocean-kubernetes