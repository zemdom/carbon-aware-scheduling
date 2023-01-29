# carbon-aware-scheduling - workloads run on Kubernetes cluster

## Prework
**1. Prepare Python environment on Ansible controller:**
```shell
python3 -m venv ./venv
source ./venv/bin/activate
pip3 install -r requirements.txt
```

How to exit Python environment:
```shell
deactivate
```

\
**2. Prepare Ansible collections on Ansible controller:**
```shell
ansible-playbook ansible-playbook/setup-collections.yaml
```

\
**3. Create inventory file:**\
To create an inventory file, copy the content of `ansible-playbook/inventory.example` to `ansible-playbook/inventory`
file and replace the example hostnames for `[k8s_nodes]` group with own hostnames or IPs.

\
**4. *[optional]* Create secret file with Docker registry credentials:**\
To push custom `benchmark-operator` Docker image to registry which requires authorization, copy the content
of `ansible-playbook/vars/env.yml.example` to `ansible-playbook/vars/env.yml` file and specify variables values.
Then, pass the new file as command-line argument when invoking playbooks, e.g.:
```shell
ansible-playbook ansible-playbook/run-all.yaml -i ansible-playbook/inventory --extra-vars "@ansible-playbook/vars/env.yaml"
```

## Execution
You can choose to run playbooks step by step or to run all the steps at once.

Running step by step:

**1. Start monitoring k8s cluster:**
```shell
ansible-playbook ansible-playbook/deploy-k8s-nodes-monitoring.yaml
```

Check if deployment is running:
```shell
kubectl get all --namespace nodes-monitoring
```

Expose kube-prometheus outside cluster:
```shell
kubectl port-forward svc/kube-prometheus-prometheus 9090:9090 --namespace nodes-monitoring
```

Check if metrics-server-exporter is running:
```shell
curl "http://localhost:9090/api/v1/query?query=kube_metrics_server_nodes_cpu"
```

Uninstall kube-prometheus chart:
```shell
ansible-playbook ansible-playbook/uninstall-k8s-nodes-monitoring.yaml
```

\
**2. Deploy benchmark-operator:**
```shell
ansible-playbook ansible-playbook/deploy-benchmark-operator.yaml
```

Build and deploy benchmark-operator:
```shell
build_benchmark_operator=True ansible-playbook ansible-playbook/deploy-benchmark-operator.yaml --extra-vars "@ansible-playbook/vars/env.yaml"
```

Check if deployment is running:
```shell
kubectl get all --namespace benchmark-operator
```

\
**3. Run first round of benchmarks:**
```shell
[pin_to_nodes=<True|False>] [run_id=<optional>] ansible-playbook ansible-playbook/run-synthetic-benchmarks.yaml -i ansible-playbook/inventory
```

Check benchmarks deployment status:
```shell
kubectl get benchmark.ripsaw.cloudbulldozer.io --namespace benchmark-operator
```

To rerun the test with the same scheduling decisions (for non-pinned testcase) execute command:
```shell
previous_run_id=<run-id-of-test-to-recreate> [run_id=<optional>] ansible-playbook ansible-playbook/run-synthetic-benchmarks.yaml -i ansible-playbook/inventory
```

\
**4. Run second round of benchmarks:**
```shell
ansible-playbook ansible-playbook/run-workload-benchmarks.yaml -i ansible-playbook/inventory
```

\
**5. Stop monitoring k8s cluster:**
```shell
ansible-playbook ansible-playbook/uninstall-k8s-nodes-monitoring.yaml
```

\
**6. Uninstall benchmark-operator:**
```shell
helm uninstall benchmark-operator --namespace benchmark-operator
```

\
All above steps are equivalent to running:
```shell
[build_benchmark_operator=True] [pin_to_nodes=<True|False>] [run_id=<optional>] ansible-playbook ansible-playbook/run-all.yaml -i ansible-playbook/inventory [--extra-vars "@ansible-playbook/vars/env.yaml"]
```