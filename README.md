# go-web-app
Create a simple Hello World  web app, in GO.

Create the Docker image and deploy on kubernetes.


## Build/Load image to kind nodes
```sh
cd go-web-app/
docker build -t adolfomaltez/go-web-app .
kind load docker-image adolfomaltez/go-web-app:latest --name cluster
```

## deploy and expose go web app
```sh
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl port-forward service/go-web-service 3000:80
```

Go to http://localhost:3000

# Reference:
- https://www.digitalocean.com/community/tutorials/how-to-deploy-resilient-go-app-digitalocean-kubernetes