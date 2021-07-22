How to deploy locally on KIND:
0. install KIND, install kustomize
1. run "kind create cluster"
2. run "kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml"
3. run "make docker-build"... you may want to delete the "test" command in the makefile for this to work, since
tests require a TLS Cert in the right place on your local, whereas actually deploying uses step 2's certmanager to do that. 
4. run "kind load docker-image controller:latest"
5. edit the manager.yaml in /config/manager and add "ImagePullPolicy: Never" in the appropriate spot
6. run "make install"
7. run "make deploy"
8. run "kubectl create -f config/samples/cache_v1alpha1_memcached.yaml -n  memcached-operator-system"

TODO: Figure out how to get the metrics fully up-and-running.