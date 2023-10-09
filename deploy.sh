#!/opt/homebrew/bin/zsh

set -ex;

export PIPELINE_ID=$(date "+%Y%m%d%H%M%S")

#build
docker buildx build -t hamstershare/computeshare-server:${PIPELINE_ID} --platform=linux/amd64 --push .

envsubst < k8s.yml | kubectl -n computeshare apply -f -
