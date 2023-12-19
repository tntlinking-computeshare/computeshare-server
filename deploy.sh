#!/opt/homebrew/bin/zsh

set -ex;

export PIPELINE_ID=$(date "+%Y%m%d%H%M%S")

#build
docker buildx build -t hamstershare/computeshare-server:${PIPELINE_ID} --platform=linux/amd64 --push .

##envsubst < k8s.yml | kubectl --kubeconfig /Users/mohaijiang/.kube/config_computeshare -n computeshare apply -f -
kubectl -n computeshare set image deployment/computeshare-server computeshare-server=hamstershare/computeshare-server:${PIPELINE_ID}
