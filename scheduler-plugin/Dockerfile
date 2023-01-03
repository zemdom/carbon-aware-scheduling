FROM golang:1.18-buster

ARG TAG

# download and install Tensorflow C library
RUN curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.9.1.tar.gz" | tar -C /usr/local -xz && ldconfig

WORKDIR /go/src/sigs.k8s.io/scheduler-plugin
COPY ./src .
RUN go build -ldflags "-X k8s.io/component-base/version.gitVersion=${TAG} -w" -o bin/kube-scheduler main.go

RUN mv /go/src/sigs.k8s.io/scheduler-plugin/bin/kube-scheduler /usr/local/bin/kube-scheduler && rm -rf /go/src/sigs.k8s.io/scheduler-plugin

CMD ["kube-scheduler"]