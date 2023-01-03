import os
import sys

from constants import USE_MEMORY_FEATURE, CREATE_PIECEWISE_LINEAR_CPU_MODEL, EPOCHS, ONLY_PREPROCESS_DATA
from models.linear_model import create_model, train_model, use_model, save_model
from utils.data_processing import load_data, extract_host_data, create_datasets


def _handle_script_invocation_parameters(arg_vector):
    script_name = arg_vector[0]
    if len(arg_vector) < 4:
        print(
            f'usage: {script_name} data-file-path models-output-path run-id '
            f'[only-preprocess-data (default: False)] '
            f'[epochs (default: 300)] '
            f'[use-memory-feature (default: True)] '
            f'[create-piecewise-linear-cpu-model (default: False)]'
        )
        sys.exit(1)

    data_file_path = arg_vector[1]
    models_output_path = arg_vector[2]
    run_id = arg_vector[3]
    only_preprocess_data = arg_vector[4] == 'True' if len(arg_vector) > 4 else ONLY_PREPROCESS_DATA
    epochs = int(arg_vector[5]) if len(arg_vector) > 5 else EPOCHS
    use_memory_feature = arg_vector[6] == 'True' if len(arg_vector) > 6 else USE_MEMORY_FEATURE
    create_piecewise_linear_cpu_model = arg_vector[7] == 'True' if len(
        arg_vector) > 7 else CREATE_PIECEWISE_LINEAR_CPU_MODEL

    return data_file_path, models_output_path, run_id, only_preprocess_data, epochs, use_memory_feature, create_piecewise_linear_cpu_model


def _generate_models_output_path(models_output_path, run_id, use_memory_feature, create_piecewise_linear_cpu_model):
    model_type = '_linear'
    features = '_cpu-memory-feature' if use_memory_feature else '_cpu-feature'
    piecewise_linear_cpu_model = f'_piecewise-{create_piecewise_linear_cpu_model}-models' if \
        create_piecewise_linear_cpu_model else ''

    models_destination_directory = f'{run_id}{model_type}{features}{piecewise_linear_cpu_model}'
    models_output_path = os.path.join(models_output_path, models_destination_directory)

    return models_output_path


def main():
    data_file_path, models_output_path, run_id, only_preprocess_data, epochs, use_memory_feature, create_piecewise_linear_cpu_model = \
        _handle_script_invocation_parameters(sys.argv)
    all_hosts_utilization_data = load_data(data_file_path)

    models_output_path = _generate_models_output_path(models_output_path, run_id, use_memory_feature,
                                                      create_piecewise_linear_cpu_model)

    for hostname, host_data in all_hosts_utilization_data.items():
        host_dataset = extract_host_data(host_data, models_output_path, hostname, use_memory_feature)

        if only_preprocess_data:
            continue

        train_features, train_labels, test_features, test_labels = create_datasets(host_dataset, models_output_path,
                                                                                   hostname,
                                                                                   create_piecewise_linear_cpu_model)

        for dataset_index in range(len(train_features)):
            dataset_train_features = train_features[dataset_index]
            dataset_train_labels = train_labels[dataset_index]
            dataset_test_features = test_features[dataset_index]
            dataset_test_labels = test_labels[dataset_index]

            model = create_model(dataset_train_features, use_memory_feature)
            model = train_model(model, dataset_train_features, dataset_train_labels, models_output_path, hostname,
                                dataset_index, epochs, create_piecewise_linear_cpu_model)
            use_model(model, dataset_test_features, dataset_test_labels, models_output_path, hostname, dataset_index,
                      create_piecewise_linear_cpu_model)
            save_model(model, models_output_path, hostname, dataset_index, create_piecewise_linear_cpu_model)


if __name__ == '__main__':
    main()
