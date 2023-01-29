package score_plugin

import (
	"context"
	"github.com/zemdom/carbon-aware-scheduling/scheduler-plugin/src/power-model"
	"github.com/zemdom/carbon-aware-scheduling/scheduler-plugin/src/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"math"
)

const (
	Name = "CarbonScorePlugin"

	// MaxNodeScore is the maximum score a Score plugin is expected to return.
	MaxNodeScore int64 = 100

	// MinNodeScore is the minimum score a Score plugin is expected to return.
	MinNodeScore int64 = 0

	// MaxTotalScore is the maximum total score.
	MaxTotalScore int64 = math.MaxInt64
)

type CarbonScorePlugin struct {
	handle        framework.Handle
	metricsClient *metrics.Clientset
	powerModels   map[string]*power_model.PowerModel
}

var _ framework.ScorePlugin = &CarbonScorePlugin{}

func New(obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	var master string
	kubeconfig := "/etc/kubernetes/scheduler.conf"
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}
	metricsClient, err := metrics.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	powerModels := make(map[string]*power_model.PowerModel)

	return &CarbonScorePlugin{
		handle:        handle,
		metricsClient: metricsClient,
		powerModels:   powerModels,
	}, nil
}

func (p *CarbonScorePlugin) Name() string {
	return Name
}

func (p *CarbonScorePlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	score := int64(0)

	usedResources := p.getNodeResources(ctx, nodeName)

	nodePowerModel := p.getNodePowerModel(nodeName)
	modelInputData := usedResources
	scoreFloat32 := nodePowerModel.Predict(modelInputData)
	score = int64(scoreFloat32 * 1000)

	klog.V(2).Infof("[CarbonScorePlugin] Node [%v] predicted power usage: %f (when scheduling Pod: [+%v])", nodeName, scoreFloat32, pod.Name)
	//klog.V(2).Infof("[CarbonScorePlugin] Pod [%v] score for node [%+v]: %d", pod.Name, nodeName, score)

	return score, framework.NewStatus(framework.Success, "")
}

func (p *CarbonScorePlugin) getNodeResources(ctx context.Context, nodeName string) []float32 {
	nodeMetrics, err := p.metricsClient.MetricsV1beta1().NodeMetricses().Get(ctx, nodeName, metav1.GetOptions{})
	//ref: https://stackoverflow.com/questions/52763291/get-current-resource-usage-of-a-pod-in-kubernetes-with-go-client
	if err != nil {
		return nil
	}

	cpu := nodeMetrics.Usage.Cpu()
	memory := nodeMetrics.Usage.Memory()

	klog.V(2).Infof("[CarbonScorePlugin] Node [%+v] used resources: CPU -> %d, memory -> %d", nodeName, cpu, memory)

	cpuFloat32 := float32(cpu.MilliValue()) / 1000
	memoryFloat32 := float32(memory.Value()) / 1024

	requests := make([]float32, 0)
	requests = append(requests, cpuFloat32)
	requests = append(requests, memoryFloat32)

	return requests
}

func (p *CarbonScorePlugin) getNodePowerModel(nodeName string) *power_model.PowerModel {
	nodePowerModel, isLoaded := p.powerModels[nodeName]
	if !isLoaded {
		nodePowerModel = power_model.New(nodeName)
		p.powerModels[nodeName] = nodePowerModel
	}

	return nodePowerModel
}

func (p *CarbonScorePlugin) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	highest := int64(0)
	lowest := int64(math.MaxInt64)
	for _, nodeScore := range scores {
		highest = utils.MaxOf(highest, nodeScore.Score)
		lowest = utils.MinOf(lowest, nodeScore.Score)
	}

	for i, nodeScore := range scores {
		score := normalizeNodeScore(nodeScore.Score, lowest, highest)
		scores[i].Score = MaxNodeScore - score // revert nodes scores - the lowest power usage should be scored the highest

		klog.V(2).Infof("[CarbonScorePlugin] Pod [%v] normalized score for node [%+v]: %d", pod.Name, scores[i].Name, scores[i].Score)
	}
	return nil
}

func normalizeNodeScore(nodeScore int64, lowest int64, highest int64) int64 {
	return (nodeScore-lowest)/(highest-lowest)*(MaxNodeScore-MinNodeScore) + MinNodeScore
}

func (p *CarbonScorePlugin) ScoreExtensions() framework.ScoreExtensions {
	return p
}
