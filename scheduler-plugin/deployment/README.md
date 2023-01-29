# carbon-aware-scheduling - deploying extended kube-scheduler

Ansible playbooks contained in this directory can be used for building and deploying custom `kube-scheduler` image -
extended with own Kubernetes Scheduling Framework plugins.

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
**2. Setup ssh keys for Ansible (when using minikube):**
```shell
cat ~/.ssh/id_rsa.pub | ssh -i ~/.minikube/machines/minikube/id_rsa docker@$(minikube ip) 'cat >> ~/.ssh/authorized_keys'
```

\
**3. Create inventory file:**\
   To create an inventory file, copy the content of `ansible-playbook/inventory.example` to `ansible-playbook/inventory`
   file and replace the example hostnames for `[control_plane]` group with own hostnames or IPs.

\
**4. *[optional]* Create secret file with Docker registry credentials:**\
   To push `kube-scheduler` Docker image to registry which requires authorization, copy the content
   of `ansible-playbook/vars/env.yml.example` to `ansible-playbook/vars/env.yml` file and specify variables values.
   Then, pass the new file as command-line argument when invoking playbooks, e.g.:
```shell
ansible-playbook ansible-playbook/run-all.yaml -i ansible-playbook/inventory --extra-vars "@ansible-playbook/vars/env.yaml"
```

## Execution
The scripts can be run step by step or all steps at once.

\
Running step by step:
**1. Build Docker image of extended scheduler:**

```shell
scheduler_plugin_image_tag="vX.Y.Z<-debug>" [debug=True|False] ansible-playbook ansible-playbook/build-scheduler-plugin.yaml -i ansible-playbook/inventory [--extra-vars "@ansible-playbook/vars/env.yaml"]
```

\
**2. Deploy extended scheduler to k8s cluster:**
```shell
scheduler_plugin_image_tag="vX.Y.Z<-debug>" [debug=True|False] ansible-playbook ansible-playbook/deploy-scheduler-plugin.yaml -i ansible-playbook/inventory
```

**3. Uninstall extended scheduler from k8s cluster:**
```shell
ansible-playbook ansible-playbook/deploy-scheduler-plugin.yaml -i ansible-playbook/inventory
```

\
Executing all steps described above at once:
```shell
scheduler_plugin_image_tag="vX.Y.Z-debug>" [debug=True|False] ansible-playbook ansible-playbook/run-all.yaml -i ansible-playbook/inventory [--extra-vars "@ansible-playbook/vars/env.yaml"]
```

### Execution - debug mode
To debug extended scheduler running in k8s cluster, run steps described above with `debug` variable value set to `True`. Then, expose the pod outside the cluster:
```shell
kubectl port-forward pod/kube-scheduler-<master-name> -n kube-system 40000:40000
# ref: https://itnext.io/debug-a-go-application-in-kubernetes-from-ide-c45ad26d8785
```
Next, add new `Run/Debug configuration` in Goland with parameters:
```shell
Go Remote
host: localhost
port: 40000
```
set breakpoints in local IDE and run the configuration.