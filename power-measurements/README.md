# carbon-scheduling - power measurements on servers forming a Kubernetes cluster

This directory contains instructions on how to measure servers power consumption using Yokogawa power meter.

## Prework
**1. Prepare Python environment on controller machine (the one the power meter is plugged into):**
```shell
python3 -m venv ./venv
source ./venv/bin/activate
```

How to exit Python environment:
```shell
deactivate
```

\
**2. Install yokotool:**
```shell
git clone https://github.com/intel/yoko-tool.git
pip install ./yoko-tool
```

\
**3. Create yokotool config file:**
```shell
sudo tee /etc/yokotool.conf > /dev/null << EOF
[default]
devnode=/dev/usbtmc0
pmtype=wt310
EOF
```

## Execution
**1. Start monitoring server(s) power usage:**
```shell
yokotool read T,P -o <run_id>.csv
```

\
**2. Run chosen workload on a measured machine.**

\
**3. Stop monitoring after finishing workload by pressing CTRL+C.**
