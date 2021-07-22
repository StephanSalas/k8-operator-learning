How to deploy locally on KIND:
1. install KIND, install kustomize
2. run "kind create cluster"
3. run "kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml"
4. run "make docker-build"... you may want to delete the "test" command in the makefile for this to work, since
tests require a TLS Cert in the right place on your local, whereas actually deploying uses step 2's certmanager to do that. 
5. run "kind load docker-image controller:latest"
6. edit the manager.yaml in /config/manager and add "ImagePullPolicy: Never" in the appropriate spot
7. run "make install"
8. run "make deploy"
9. run "kubectl create -f config/samples/v1alpha1_application.yaml -n  application-operator-system"

TODO: Figure out how to get the metrics fully up-and-running.
