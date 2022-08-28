# Istio Zero Trust Demo App

## Prerequisites

### CLIs/Tools

* [kubectl](https://kubernetes.io/docs/tasks/tools/)
* [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/)
* [istioctl](https://istio.io/latest/docs/reference/commands/istioctl/)
* [Rancher Desktop](https://docs.rancherdesktop.io/getting-started/installation)

### Deployed in Cluster

* [KubeWarden Operator](https://docs.kubewarden.io/quick-start) installed (requires cert-manager).
* [Istio Operator](https://istio.io/latest/docs/setup/install/operator/) is installed onto your cluster.
* An IstioOperator custom resource is created. There is an [example](https://github.com/atoy3731/istio-demo-app/blob/main/k8s/examples/istiooperator.yaml) here.
* A cert-manager [Certificate](https://cert-manager.io/docs/usage/) created in the istio-system namespace that matches the below gateway configuration.
* A Gateway custom resource is created. There is an [example](https://github.com/atoy3731/istio-demo-app/blob/main/k8s/examples/gateway.yaml) here.

## Setting up '/etc/hosts'

Since we're assuming real DNS isn't available, we need to configure '/etc/hosts' to spoof it. Run the following commands to get the IP:

```bash
kubectl get svc -n istio-system istio-ingressgateway -o jsonpath={.status.loadBalancer.ingress[0].ip}
```

Now using `sudo`, update your `/etc/hosts` file and add the following line to the bottom:
```bash
IP_FROM_ABOVE app1.example.com
```

## Deploying the base applications

With your kube context set for your desired cluster, run the following at the root of this repo:

```bash
kustomize build k8s/base/ | kubectl apply -f -
```

Wait for the pods to be healthy and 2/2:

```bash
kubectl get pods -n app1 -w
```

Once they are, navigate to https://app1.example.com - you may need to close/reopen your browser if the IP is cached.

## Layering on Zero-Trust

Go to the `k8s/working` directory and start applying files:

* *1-peer-authentication.yaml*: Enforces that only istio-injected pods can talk to other istio-injected pods. Denies traffic to/from non-istio-injected pods.

* *2-deny-authpol.yaml*: Sets default-deny ingress rule for all istio-injected pods across the entire mesh.

* *3-allow-istio-ingress.yaml*: Allow traffic into demo pods in `app1` namespace explicitly from istio-ingressgateway pods.

* *4-allow-intra-namespace.yaml*: Allow traffic from all pods in `app1` namespace to other pods in `app1` namespace.

* *5-allow-cross-namespace.yaml*: Allow traffic from pods in `app1` namespace to demo pods in `app2` namespace, only on port 8080 when a `GET` to `/status` is used. Examples of `when` clauses as well.

* *Setting REGISTRY_ONLY*: To set `REGISTRY_ONLY` globally, modify your IstioOperator custom resource and change the `.meshConfig.outboundTrafficPolicy.mode` from `ALLOW_ALL` to `REGISTRY_ONLY`.

* *6-rancher-service-entry.yaml*: Allow traffic to `www.rancher.com` from pods in the `app1` namespace. Important to note the `exportTo`, since by default service entries are globally applied.

* *7-kubewarden-istio-policy.yaml*: Enforces namespaces and pods cannot have istio injection disabled from webhook policy enforcement.


