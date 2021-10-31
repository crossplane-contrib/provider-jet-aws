# Testing CRD Creation

## Install monitoring stack

```bash
# Details in https://prometheus-operator.dev/docs/prologue/quick-start/
git clone https://github.com/prometheus-operator/kube-prometheus.git && cd kube-prometheus
kubectl create -f manifests/setup
until kubectl get servicemonitors --all-namespaces ; do date; sleep 1; echo ""; done
kubectl create -f manifests/
```

### Patch Prometheus

Run the following command to patch prometheus with `prometheus.yaml` file which
has 10s interval.

```bash
# Change content of prometheus.yaml as you'd like.
echo -n '{"data":{"prometheus.yaml.gz": "' > /tmp/patch.json && gzip < prometheus.yaml > prometheus.yaml.gz && cat prometheus.yaml.gz | base64 -w 0 >> /tmp/patch.json && echo '"}}' >> /tmp/patch.json
kubectl patch secret -n monitoring prometheus-k8s --patch "$(cat /tmp/patch.json)"
```

To get the current settings:
```bash
kubectl get secret -n monitoring prometheus-k8s -o jsonpath='{.data.prometheus\.yaml\.gz}' | base64 -d > prometheus.yaml.gz && gunzip -c prometheus.yaml.gz > prometheus.yaml
```

### Patch Grafana Data Source

Grafana data source patch (to make interval 10s):
```bash
echo -n '{"data":{"datasources.yaml": "' > /tmp/patch.json && cat grafana.json | base64 -w 0 >> /tmp/patch.json && echo -n '"}}' >> /tmp/patch.json
kubectl patch secret -n monitoring grafana-datasources --patch "$(cat /tmp/patch.json)"
```

Get current Grafana data source settings:
```bash
kubectl get secret -n monitoring grafana-datasources -o jsonpath='{.data.datasources\.yaml}' | base64 -d > grafana-current.json
```


Login to Grafana with `admin/admin`:
```bash
kubectl --namespace monitoring port-forward svc/grafana 4000:3000
```

## Build and Use Go Program as a Job

Build the binary and image:
```bash
IMAGE_PATH="muvaf/scale-test:v1.3"
cd $GOPATH/crossplane/provider-tf-aws
GOOS=linux GOARCH=amd64 go build -o cmd/test/main cmd/test/main.go && chmod +x cmd/test/main

docker build . -f cmd/test/Dockerfile -t $IMAGE_PATH
docker push $IMAGE_PATH
```

Deploy the `Job` to start:
```bash
kubectl apply -f $GOPATH/crossplane/provider-tf-aws/cmd/test/job.yaml
```