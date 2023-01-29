package main

import (
	"github.com/zemdom/carbon-aware-scheduling/scheduler-plugin/src/score-plugin"
	"k8s.io/component-base/cli"
	"os"
)

import (
	scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	command := scheduler.NewSchedulerCommand(
		scheduler.WithPlugin(score_plugin.Name, score_plugin.New),
	)

	code := cli.Run(command)
	os.Exit(code)
}
