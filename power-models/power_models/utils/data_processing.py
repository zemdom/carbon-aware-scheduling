import json
import os

import numpy as np
import pandas as pd

from constants import LABEL_NAME, PIECEWISE_LINEAR_CPU_MODELS_COUNT, DATASETS_SPLIT
from utils.plots import create_dataset_statistics_plots
from utils.utils import create_directory


def load_data(data_file_path):
    # get all files *_node_resources_info.json
    with open(data_file_path) as input_file:
        utilization_data = json.load(input_file)

    return utilization_data


def extract_host_data(host_data, models_output_path, hostname, use_memory_feature):
    cpu_utilization = []
    memory_utilization = []
    power_draw = []

    for timestamp_data in host_data.values():
        if len(timestamp_data) != 2:  # drop incomplete entries
            # if len(timestamp_data) != 3:  # drop incomplete entries
            continue

        cpu = timestamp_data[0] / 1000000000  # convert nanoCPUs to CPUs
        memory = timestamp_data[1]
        # power = timestamp_data[2]
        power = 3000000.0 * cpu + 1.0 * memory + 1

        cpu_utilization.append(cpu)
        memory_utilization.append(memory)
        power_draw.append(power)

    cpu_utilization = np.array(cpu_utilization, dtype=float)
    memory_utilization = np.array(memory_utilization, dtype=float)
    power_draw = np.array(power_draw, dtype=float)

    if use_memory_feature:
        host_dataset = pd.DataFrame({
            'cpu': cpu_utilization,
            'memory': memory_utilization,
            LABEL_NAME: power_draw
        })
    else:
        host_dataset = pd.DataFrame({
            'cpu': cpu_utilization,
            LABEL_NAME: power_draw
        })

    timestamps = list(map(int, list(host_data.keys())))
    create_dataset_statistics_plots(timestamps, cpu_utilization, memory_utilization, power_draw, host_dataset,
                                    models_output_path, hostname, use_memory_feature)

    return host_dataset


def create_datasets(original_dataset, models_output_path, hostname, create_piecewise_linear_cpu_models):
    datasets = []

    if create_piecewise_linear_cpu_models:
        cpu_utilization_thresholds = _calculate_cpu_utilization_thresholds(original_dataset)
        _save_cpu_utilization_thresholds_to_file(cpu_utilization_thresholds, models_output_path, hostname)

        for threshold_index in range(0, PIECEWISE_LINEAR_CPU_MODELS_COUNT):
            if threshold_index == 0:
                filter_condition = original_dataset['cpu'] <= cpu_utilization_thresholds[threshold_index]
            elif threshold_index == (PIECEWISE_LINEAR_CPU_MODELS_COUNT - 1):
                filter_condition = cpu_utilization_thresholds[threshold_index - 1] < original_dataset['cpu']
            else:
                filter_condition = (cpu_utilization_thresholds[threshold_index - 1] < original_dataset['cpu']) & (
                        original_dataset['cpu'] <= cpu_utilization_thresholds[threshold_index])

            filtered_features = original_dataset[filter_condition]
            datasets.append(filtered_features)
    else:
        datasets.append(original_dataset)

    train_features, train_labels, test_features, test_labels = [], [], [], []

    for dataset in datasets:
        train_dataset = dataset.sample(frac=DATASETS_SPLIT, random_state=0)
        test_dataset = dataset.drop(train_dataset.index)

        dataset_train_features = train_dataset.copy()
        dataset_train_labels = dataset_train_features.pop(LABEL_NAME)

        dataset_test_features = test_dataset.copy()
        dataset_test_labels = dataset_test_features.pop(LABEL_NAME)

        train_features.append(dataset_train_features)
        train_labels.append(dataset_train_labels)
        test_features.append(dataset_test_features)
        test_labels.append(dataset_test_labels)

    return train_features, train_labels, test_features, test_labels


def _calculate_cpu_utilization_thresholds(dataset):
    max_cpu_utilization = dataset['cpu'].max()
    cpu_utilization_thresholds = []

    # range -> do not specify min and max value - instead use open intervals on the ends
    for multiplier in range(1, PIECEWISE_LINEAR_CPU_MODELS_COUNT):
        cpu_utilization_threshold = multiplier * (max_cpu_utilization / PIECEWISE_LINEAR_CPU_MODELS_COUNT)
        cpu_utilization_thresholds.append(cpu_utilization_threshold)

    return cpu_utilization_thresholds


def _save_cpu_utilization_thresholds_to_file(cpu_utilization_thresholds, models_output_path, hostname):
    cpu_utilization_thresholds_string = '('

    for cpu_utilization_threshold in cpu_utilization_thresholds:
        rounded_cpu_utilization_threshold = round(cpu_utilization_threshold, 2)
        cpu_utilization_thresholds_string = f'{cpu_utilization_thresholds_string}{rounded_cpu_utilization_threshold},'

    cpu_utilization_thresholds_string = cpu_utilization_thresholds_string[:-1]  # remove redundant comma
    cpu_utilization_thresholds_string = f'{cpu_utilization_thresholds_string})'

    model_full_output_path = os.path.join(models_output_path, hostname)
    create_directory(model_full_output_path)
    with open(os.path.join(model_full_output_path, f'{hostname}_cpu-utilization-thresholds-info.log'), 'w') \
            as output_file:
        output_file.write(cpu_utilization_thresholds_string)
