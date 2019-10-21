
# create deployment
kubectl delete -f gateway/deployment.yaml
kubectl delete -f booking/deployment.yaml
kubectl delete -f tracking/deployment.yaml

# deploy services
kubectl delete -f gateway/service.yaml
kubectl delete -f booking/service.yaml
kubectl delete -f tracking/service.yaml

# create role and service account
kubectl delete -f gateway/role-binding.yaml
kubectl delete -f gateway/service-account.yaml
kubectl delete -f gateway/namespace-role.yaml