import json
import os
import sys
from datetime import datetime

import requests

# can be adjusted by script invocation parameters
DEFAULT_TIME_RESOLUTION = 60
DEFAULT_PROMETHEUS_ENDPOINT = 'http://localhost:9090'

# constant values
PROMETHEUS_QUERY_ENDPOINT = '/api/v1/query'
NODE_RESOURCES_INFO_METRIC = [
    'kube_metrics_server_nodes_cpu',  # returned value in nano CPUs
    'kube_metrics_server_nodes_mem'   # returned value in KiB
]

TARGET_DATA_FORMAT = {
    'node':
        {
            'timestamp': ['cpu (nano CPUs)', 'memory (KiB)']
        }
}


def handle_script_invocation_parameters(arg_vector):
    script_name = arg_vector[0]
    if len(arg_vector) < 3:
        print(
            f'usage: {script_name} run-id output-directory-path [scrape-time-resolution (no of minutes, default: '
            f'60)] [prometheus-endpoint (default: http://localhost:9090)]'
        )
        sys.exit(1)

    run_id = arg_vector[1]
    output_directory_path = arg_vector[2]

    time_resolution = arg_vector[3] if len(arg_vector) > 3 else DEFAULT_TIME_RESOLUTION
    time_resolution = f'[{time_resolution}m:]'

    prometheus_endpoint = arg_vector[4] if len(arg_vector) > 4 else DEFAULT_PROMETHEUS_ENDPOINT

    return run_id, output_directory_path, time_resolution, prometheus_endpoint


def get_node_resources_info(time_resolution, prometheus_endpoint):
    prometheus_query_endpoint = f'{prometheus_endpoint}{PROMETHEUS_QUERY_ENDPOINT}'

    results = []
    for metric in NODE_RESOURCES_INFO_METRIC:
        response = requests.get(prometheus_query_endpoint,
                                params={'query': f'{metric}{time_resolution}'})

        metric_results = response.json()['data']['result']
        results.append(metric_results)

        # print(f'Header: {columns_to_extract + ["timestamp", "value"]}')

    return results


def reformat_node_resources_info_into_dictionary(node_resources_info):
    results_dictionary = {}

    # merge rows with data for the same node
    for metric_info in node_resources_info:
        for node_metric_info in metric_info:
            node_name = node_metric_info['metric']['exported_instance']
            resources_info = node_metric_info['values']

            if node_name not in results_dictionary.keys():
                results_dictionary[node_name] = []
            results_dictionary[node_name].extend(resources_info)

    # flatten node resources samples -> one timestamp for multiple resources
    for node_name, resources_info in results_dictionary.items():
        node_resources_dictionary = {}
        for sample in resources_info:
            timestamp = sample[0]
            resource_value = sample[1]
            resource_value = round_node_resources_info_values(resource_value)
            if timestamp not in node_resources_dictionary:
                node_resources_dictionary[timestamp] = []
            node_resources_dictionary[timestamp].append(resource_value)
        results_dictionary[node_name] = node_resources_dictionary

    return results_dictionary


def round_node_resources_info_values(resource_value):
    if resource_value.isdigit():
        rounded_resource_value = int(resource_value)
    else:
        rounded_resource_value = float(resource_value)
        rounded_resource_value = round(rounded_resource_value, 2)

    return rounded_resource_value


def save_node_resources_info_into_file(node_resources_info_dictionary, output_directory_path, run_id, time_resolution):
    output_path = os.path.join(output_directory_path, f'{run_id}_prometheus_data')
    os.makedirs(output_path, exist_ok=True)

    timestamp = datetime.now().strftime('%Y-%m-%d_%H:%M:%S')

    with open(os.path.join(output_path, f'{run_id}_prometheus_scraping_timestamp.txt'), 'w') as timestamp_file:
        timestamp_file.write(f'File saved timestamp: {timestamp}')
        timestamp_file.write('\n')
        timestamp_file.write(f'Scraping interval: {time_resolution}')

    # with open(os.path.join(output_path, f'{run_id}_{timestamp}_node_resources_info.json'), 'w') as output_file:
    with open(os.path.join(output_path, f'{run_id}_node_resources_info.json'), 'w') as output_file:
        output_file.write(json.dumps(node_resources_info_dictionary))

    # with open(os.path.join(output_path, f'{run_id}_{timestamp}_data_format.json'), 'w') as schema_file:
    with open(os.path.join(output_path, f'{run_id}_data_format.json'), 'w') as schema_file:
        schema_file.write(json.dumps(TARGET_DATA_FORMAT))


def main():
    run_id, output_directory_path, time_resolution, prometheus_endpoint = handle_script_invocation_parameters(sys.argv)

    node_resources_info = get_node_resources_info(time_resolution, prometheus_endpoint)
    node_resources_info_dictionary = reformat_node_resources_info_into_dictionary(
        node_resources_info)
    save_node_resources_info_into_file(node_resources_info_dictionary, output_directory_path, run_id, time_resolution)


if __name__ == '__main__':
    main()
