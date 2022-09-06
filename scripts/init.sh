#!/bin/bash

kubectl cluster-info

echo ""
printf "Roll back your cluster? (y/n) "
read ROLLBACK

if [[ "$ROLLBACK" != "y" ]]; then
    echo "Cancelling"
    exit 0
fi

echo "Installing Istio.."
istioctl operator init
kubectl apply -f ../k8s/examples/istiooperator.yaml
sleep 5
kubectl apply -f ../k8s/examples/gateway.yaml

echo "Installing cert-manager.."
kubectl apply -f https://github.com/jetstack/cert-manager/releases/latest/download/cert-manager.yaml
sleep 30

echo "Installing Kubewarden.."
helm repo add kubewarden https://charts.kubewarden.io
helm install --wait -n kubewarden --create-namespace kubewarden-crds kubewarden/kubewarden-crds
helm install --wait -n kubewarden kubewarden-controller kubewarden/kubewarden-controller
helm install --wait -n kubewarden kubewarden-defaults kubewarden/kubewarden-defaults
