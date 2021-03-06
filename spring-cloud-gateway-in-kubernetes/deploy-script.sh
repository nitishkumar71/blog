# apply miniube docker image
eval $(minikube docker-env)
#build images
docker build --tag=gateway:latest gateway/.
docker build --tag=booking:latest booking/.
docker build --tag=tracking:latest tracking/.

# https://kind.sigs.k8s.io/docs/user/quick-start/
# if you are using kind k8 cluster
# kind load docker-image gateway:latest --name gateway
# kind load docker-image booking:latest --name gateway
# kind load docker-image tracking:latest --name gateway

# create role and service account
kubectl apply -f gateway/service-account.yaml
kubectl apply -f gateway/namespace-role.yaml
kubectl apply -f gateway/role-binding.yaml

# create deployment
kubectl apply -f gateway/deployment.yaml
kubectl apply -f booking/deployment.yaml
kubectl apply -f tracking/deployment.yaml

# deploy services
kubectl apply -f gateway/service.yaml
kubectl apply -f booking/service.yaml
kubectl apply -f tracking/service.yaml