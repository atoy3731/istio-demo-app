#!/bin/bash

kubectl cluster-info

echo ""
printf "Roll back your cluster? (y/n) "
read ROLLBACK

if [[ "$ROLLBACK" != "y" ]]; then
    echo "Cancelling"
    exit 0
fi

echo "Rolling back.."
kubectl delete -f ../k8s/working/
kubectl patch istiooperator -n istio-system controlplane --type=merge -p '{"spec":{"meshConfig":{"outboundTrafficPolicy":{"mode":"ALLOW_ANY"}}}}'

kubectl patch namespace app1 --type=merge -p '{"metadata":{"labels":{"istio-injection":"enabled"}}}'
kubectl rollout restart -n app1 deploy/demo