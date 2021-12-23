### deploy httpserver
```
kubectl create ns securesvc
kubectl label ns securesvc istio-injection=enabled
kubectl create -f httpserver.yaml -n securesvc
```
```
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt
# kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt
kubectl apply -f cncamp-credential.yaml -n istio-system
kubectl apply -f istio-specs.yaml -n securesvc
```

### access the httpserver via ingress
```
INGRESS_IP=$(kubectl get svc -n istio-system | grep istio-ingressgateway | awk '{print $3}');curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/healthz -v -k
```
![https-gw-healthz](./../_img/https-gw-healthz.png)