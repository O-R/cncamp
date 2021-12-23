### deploy simple
```
kubectl create ns simple
kubectl create -f simple.yaml -n simple
kubectl create -f istio-specs.yaml -n simple
```

### access the simple via ingress
```
INGRESS_IP=$(kubectl get svc -n istio-system | grep istio-ingressgateway | awk '{print $3}');curl -H "Host: simple.cncamp.io" $INGRESS_IP/hello -v
```