# k8s-secret-ui (Kubernetes UI)

![logo.png](logo.png)

Based on https://github.com/viveksinghggits/kuui

The main purpose of this application is to have a simple UI that can be used to manage the configmaps/secrets of your Kubernetes cluster.

# Installation

To use this project either you can build it from source or download the binaries.

## Build it from source

To build the project from source, please clone it on your machine using below command

```
git clone https://github.com/mheers/k8s-secret-ui.git
```

and you can build the project using the command `go build -o k8s-secret-ui`. Optionally you can move the created binary
into your path so that you can use this from where ever you want.

### Running it

To run the project you just have to execute the `k8s-secret-ui` binary by providing it the kubeconfig file. Like below
```
./k8s-secret-ui --kubeconfig=$HOME/.kube/config
# or whereever you kubeconfig file is
```

## RoadMap

- [x] Get the service deployed on Kubernetes
- [x] Configure allowed namespaces, configmaps and secrets also using regex
- [x] Allow secrets to be modified
