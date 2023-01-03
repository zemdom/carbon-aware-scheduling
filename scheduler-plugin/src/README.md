# carbon-scheduling - custom Kubernetes Scheduling Framework plugins

This directory contains source code of `kube-scheduler` plugins implemented in Golang.

## Prework - setup local development environment

1. Download and install Golang:

```shell
sudo apt update
curl -LO https://go.dev/dl/go1.18beta1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.18beta1.linux-amd64.tar.gz
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
#export PATH=$PATH:/usr/local/go/bin
```

## Execution

Refresh `go.mod` file, so it contains only dependencies needed in source code:

```shell
go mod tidy
```

Download dependencies to local directory (to prevent losing packages when moved or renamed):

```shell
go mod vendor
```