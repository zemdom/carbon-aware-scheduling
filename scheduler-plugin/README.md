# carbon-scheduling - kube-scheduler plugin

This directory contains source code of `kube-scheduler` plugins implemented in Golang and scripts, manifests and
Dockerfiles needed to deploy them to Kubernetes cluster.

## Execution

To build and deploy extended `kube-scheduler` follow instructions from [`deployment`](./deployment) subdirectory.

## Links and notes

* [k8s scheduling framework - custom scheduler plugins](https://github.com/kubernetes/enhancements/blob/master/keps/sig-scheduling/624-scheduling-framework/README.md#custom-scheduler-plugins-out-of-tree)
* [cockroachlabs - sample implementation](https://github.com/cockroachlabs/crl-scheduler)
* [cockroachlabs - sample implementation article](https://kubernetes.io/blog/2020/12/21/writing-crl-scheduler/)
* [angao - sample implementation](https://github.com/angao/scheduler-framework-sample)
* https://medium.com/@juliorenner123/k8s-creating-a-kube-scheduler-plugin-8a826c486a1
* https://github.com/juliorenner/scheduler-plugins/blob/master/hack/update-codegen.sh
* https://github.com/kubernetes-sigs/scheduler-plugins/blob/master/doc/develop.md
* https://developer.ibm.com/articles/creating-a-custom-kube-scheduler/
* https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-scheduler/app/server_test.go
* https://github.com/kubernetes/kubernetes/blob/master/cmd/kube-scheduler/scheduler.go