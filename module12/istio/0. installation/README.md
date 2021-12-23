# Installation

## install istio

```
# curl -L https://istio.io/downloadIstio | sh -
curl -fsSL -O https://github.com/istio/istio/releases/download/1.12.1/istio-1.12.1-linux-amd64.tar.gz
cd istio-1.12.1
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```

## istio monitoring

```
grafana dashboard 7639
```