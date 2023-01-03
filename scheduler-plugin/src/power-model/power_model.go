package power_model

import (
	"fmt"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	tg "github.com/galeone/tfgo"
	"github.com/zemdom/carbon-scheduling/scheduler-plugin/src/utils"
	"log"
	"path/filepath"
	"strconv"
)

const (
	ModelDirectoryPath                   string = "/etc/power_models"
	InputTensorInfoFilename              string = "_input-tensor-info.log"
	CpuUtilizationThresholdsInfoFilename string = "_cpu-utilization-thresholds-info.log"
)

type PowerModel struct {
	models                        []*tg.Model
	inputTensorShape              []int64
	cpuUtilizationThresholds      []float32
	piecewiseLinearCpuModelsCount int
}

func New(nodeName string) *PowerModel {
	cpuUtilizationThresholds := loadCpuUtilizationThresholdsInfoFromFile(nodeName)

	var piecewiseLinearCpuModelsCount int
	if cpuUtilizationThresholds == nil {
		piecewiseLinearCpuModelsCount = 0
	} else {
		piecewiseLinearCpuModelsCount = len(cpuUtilizationThresholds) + 1
	}

	models := loadTfModels(nodeName, piecewiseLinearCpuModelsCount)
	inputTensorShape := loadInputTensorInfoFromFile(nodeName, piecewiseLinearCpuModelsCount)

	return &PowerModel{
		models:                        models,
		inputTensorShape:              inputTensorShape,
		cpuUtilizationThresholds:      cpuUtilizationThresholds,
		piecewiseLinearCpuModelsCount: piecewiseLinearCpuModelsCount,
	}
}

func loadCpuUtilizationThresholdsInfoFromFile(nodeName string) []float32 {
	cpuUtilizationThresholdsInfoFilename := fmt.Sprintf("%s%s", nodeName, CpuUtilizationThresholdsInfoFilename)
	cpuUtilizationThresholdsInfoPath := filepath.Join(ModelDirectoryPath, nodeName, cpuUtilizationThresholdsInfoFilename)

	lineContent, err := utils.LoadTupleFromFile(cpuUtilizationThresholdsInfoPath)
	if err != nil {
		return nil
	}

	cpuUtilizationThresholds := make([]float32, 0)
	for _, element := range lineContent {
		cpuUtilization, err := strconv.ParseFloat(element, 32)
		if err != nil {
			log.Fatal(err)
		}
		cpuUtilizationFloat32 := float32(cpuUtilization)
		cpuUtilizationThresholds = append(cpuUtilizationThresholds, cpuUtilizationFloat32)
	}

	return cpuUtilizationThresholds
}

func loadTfModels(nodeName string, piecewiseLinearCpuModelsCount int) []*tg.Model {
	models := make([]*tg.Model, 0)
	if piecewiseLinearCpuModelsCount == 0 {
		modelName := fmt.Sprintf("%s_model", nodeName)
		modelPath := filepath.Join(ModelDirectoryPath, nodeName, modelName)

		model := tg.LoadModel(modelPath, []string{"serve"}, nil)

		models = append(models, model)
	} else {
		var modelName, modelPath string
		var model *tg.Model

		for modelIndex := 0; modelIndex < piecewiseLinearCpuModelsCount; modelIndex++ {
			modelName = fmt.Sprintf("%s_model%d-of-%d", nodeName, modelIndex+1, piecewiseLinearCpuModelsCount)
			modelPath = filepath.Join(ModelDirectoryPath, nodeName, modelName)

			model = tg.LoadModel(modelPath, []string{"serve"}, nil)
			models = append(models, model)
		}
	}

	return models
}

func loadInputTensorInfoFromFile(nodeName string, piecewiseLinearCpuModelsCount int) []int64 {
	var inputTensorInfoFilename string

	if piecewiseLinearCpuModelsCount == 0 {
		inputTensorInfoFilename = fmt.Sprintf("%s_model%s", nodeName, InputTensorInfoFilename)
	} else {
		// piecewise linear models have the same input tensor shape - it is sufficient to load one
		inputTensorInfoFilename = fmt.Sprintf("%s_model1-of-%d%s", nodeName, piecewiseLinearCpuModelsCount, InputTensorInfoFilename)
	}

	inputTensorInfoPath := filepath.Join(ModelDirectoryPath, nodeName, inputTensorInfoFilename)
	lineContent, _ := utils.LoadTupleFromFile(inputTensorInfoPath)

	inputTensorShape := make([]int64, 0)
	for _, element := range lineContent {
		shapeDimension, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		inputTensorShape = append(inputTensorShape, shapeDimension)
	}

	return inputTensorShape
}

func (p *PowerModel) GetInputTensorShape() []int64 {
	return p.inputTensorShape
}

func (p *PowerModel) Predict(data []float32) float32 {
	inputTensor := p.float32DataToTensor(data)
	cpuUtilization := data[0]
	prediction := p.makePrediction(inputTensor, cpuUtilization)

	return prediction
}

func (p *PowerModel) float32DataToTensor(data []float32) *tf.Tensor {
	tenorData := make([][]float32, p.inputTensorShape[0])
	for index := range tenorData {
		tenorData[index] = make([]float32, p.inputTensorShape[1])
	}

	for index, element := range data {
		elementFloat32 := element
		// input tensor should always be 1D array with changing number of elements
		tenorData[0][index] = elementFloat32
	}

	tensor, _ := tf.NewTensor(tenorData)

	return tensor
}

func (p *PowerModel) makePrediction(inputTensor *tf.Tensor, cpuUtilization float32) float32 {
	modelIndex := 0

	if p.piecewiseLinearCpuModelsCount != 0 {
		modelIndex = p.piecewiseLinearCpuModelsCount - 1
		for index, cpuUtilizationThreshold := range p.cpuUtilizationThresholds {
			if cpuUtilization <= cpuUtilizationThreshold {
				modelIndex = index
				break
			}
		}
	}

	results := p.models[modelIndex].Exec([]tf.Output{
		p.models[modelIndex].Op("StatefulPartitionedCall", 0),
	}, map[tf.Output]*tf.Tensor{
		p.models[modelIndex].Op("serving_default_normalization_input", 0): inputTensor,
	})

	predictionValue := results[0].Value()
	prediction := predictionValue.([][]float32)[0][0]

	return prediction
}
