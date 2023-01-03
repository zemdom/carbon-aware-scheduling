# carbon-scheduling - cluster carbon intensity blockchain database


## Prework

### Prepare k8s cluster
To create a production blockchain network, ensure you have the following:
* One running Kubernetes cluster and the config file (`kubeconfig.yaml`) per Organization;
* One running Hashicorp Vault server per Organization (unsealed and configured);

> reference: [Hyperledger Bevel - Setting up a DLT/Blockchain network](https://blockchain-automation-framework.readthedocs.io/en/latest/operations/setting_dlt.html)

\
**1. Start one node minikube cluster:**
```console
[root@ansible-controller]# minikube start
```

\
**2. Install Hashicorp Vault on minikube cluster:**
```console
[root@ansible-controller]# cd ~/carbon-scheduling/immutable-database

# add hashicorp helm repo and install vault in default namespace
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install vault hashicorp/vault --values helm-vault-raft-values.yml --version 0.22.0

# initialize root key
kubectl exec vault-0 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json

# unseal key
sudo apt install -y jq
jq -r ".unseal_keys_b64[]" cluster-keys.json
VAULT_UNSEAL_KEY=$(jq -r ".unseal_keys_b64[]" cluster-keys.json)

# unseal vault-0
kubectl exec vault-0 -- vault operator unseal $VAULT_UNSEAL_KEY

# join Vault pods to Raft cluster
kubectl exec -ti vault-1 -- vault operator raft join http://vault-0.vault-internal:8200
kubectl exec -ti vault-2 -- vault operator raft join http://vault-0.vault-internal:8200

# unseal remaining Vaults
kubectl exec -ti vault-1 -- vault operator unseal $VAULT_UNSEAL_KEY
kubectl exec -ti vault-2 -- vault operator unseal $VAULT_UNSEAL_KEY
```
 > reference: [Hashicorp - Vault Installation to Minikube via Helm](https://learn.hashicorp.com/tutorials/vault/kubernetes-minikube-raft?in=vault/kubernetes)

\
**3. Destroy minikube cluster:**
```console
[root@ansible-controller]# minikube delete --all
```

### Prepare Ansible controller
**1. Configure Ansible controller:**
```console
[root@ansible-controller]# cd ~/carbon-scheduling/immutable-database

python3 -m venv ./venv
source ./venv/bin/activate

pip3 install -r requirements.txt
```

\
**2. Install Node.js (required to use flux GitOps tool):**
```console
[root@ansible-controller]# curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
[root@ansible-controller]# sudo apt-get install -y nodejs
```

## Execution
**1. Fork or import [Bevel GitHub repo](https://github.com/hyperledger/bevel.git) to own GitHub repository.**

\
**2. Generate GitHub repo access token.**

\
**3. Generate ssh key for k8s cluster (one key is needed for each Organization):**
```console
[root@ansible-controller:~/carbon-scheduling/immutable-database]# ssh-keygen -q -N "" -f ./gitops
```

\
**4. Get Vault root_token value:**
```console
[root@ansible-controller:~/carbon-scheduling/immutable-database]# grep "root_token" cluster-keys.json | awk -F "\"" '{print $4}'
```

\
**5. Replace variables below in network configuration file (`network-fabricv2.yaml`) with own values:**
```yaml
{{ path_to_bevel_repository }}

{{ root_token }}

{{ github_email }}
{{ github_user }}
{{ repo_access_token }}
```

\
**6. Clone Hyperledger Bevel repository fork:**
```console
[root@ansible-controller]# cd ~/carbon-scheduling/immutable-database

git clone https://github.com/<user>/bevel.git
cd bevel
git checkout b6892c031d49a1145d5ef91fe78f23e4bc3e1165
```

\
**7. Install additional requirements:**
```console
[root@ansible-controller:~/carbon-scheduling/immutable-database/bevel]# ansible-galaxy install -r platforms/shared/configuration/requirements.yaml
```

\
**8. Copy configuration files to bevel `build` directory:**
```console
[root@ansible-controller]# cd ~/carbon-scheduling/immutable-database/bevel

mkdir build

cp ~/.kube/config build/kubeconfig.yaml

cp ../network-fabricv2.yaml build/network.yaml
cp ../gitops build
cp ../inventory .
```

\
**9. Make Vault service accessible to Ansible controller:**
```console
[root@ansible-controller]# export VAULT_ADDR=http://vault:8200
[root@ansible-controller]# kubectl port-forward svc/vault 8200:8200
```

\
**10. Deploy blockchain network to k8s cluster:**
```console
[root@ansible-controller:~/carbon-scheduling/immutable-database/bevel]# ANSIBLE_PYTHON_INTERPRETER=~/carbon-scheduling/immutable-database/venv/bin/python3 ansible-playbook platforms/shared/configuration/site.yaml -e "@./build/network.yaml" -i inventory
```

\
**11. Uninstall deployment:**
```console
[root@ansible-controller]# flux uninstall
```