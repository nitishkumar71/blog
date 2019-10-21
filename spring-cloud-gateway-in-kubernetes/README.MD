## Start Minikube
```
minikub start
eval $(minikube docker-env)
```
sudo chown $USER:$USER /home/$USER/.docker -R
sudo chmod g+rwx "/home/$USER/.docker" -R

# check if gateway pod has access to read
curl -v --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt -H "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" https://kubernetes.default.svc/api/v1/namespaces/default/pods

## Build Docker Images

### API Gateway
```
docker build --tag=gateway:latest gateway/.
```
### Booking Application
```
docker build --tag=booking:latest booking/.
```

### Tracking Application
```
docker build --tag=tracking:latest tracking/.
```

### Create Service Account and Roles
```                
kubectl apply -f gateway/service-account.yaml
kubectl apply -f gateway/namespace-role.yaml
kubectl apply -f gateway/role-binding.yaml
```


### Deploy Application
```
kubectl apply -f gateway/deployment.yaml
kubectl apply -f booking/deployment.yaml
kubectl apply -f tracking/deployment.yaml
```

### Deploy Services
```
kubectl apply -f gateway/service.yaml
kubectl apply -f booking/service.yaml
kubectl apply -f tracking/service.yaml
```
