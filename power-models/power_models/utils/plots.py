from matplotlib import pyplot as plt

import datetime
import matplotlib.dates
import seaborn as sns
import os

from constants import LABEL_NAME, CPU_FEATURE_NAME, CPU_FEATURE_UNIT, LABEL_UNIT, MEMORY_FEATURE_NAME, \
    MEMORY_FEATURE_UNIT
from utils.utils import create_directory


def create_dataset_statistics_plots(timestamps, cpu_utilization, memory_utilization, power_draw, dataset,
                                    models_output_path, hostname,
                                    use_memory_feature):
    _plot_timeseries_data(timestamps, cpu_utilization, CPU_FEATURE_NAME, CPU_FEATURE_UNIT, models_output_path, hostname)
    _plot_input_data(cpu_utilization, CPU_FEATURE_NAME, CPU_FEATURE_UNIT, power_draw, LABEL_NAME, LABEL_UNIT,
                     models_output_path, hostname)

    if use_memory_feature:
        _plot_timeseries_data(timestamps, memory_utilization, MEMORY_FEATURE_NAME, MEMORY_FEATURE_UNIT,
                              models_output_path, hostname)
        _plot_input_data(memory_utilization, MEMORY_FEATURE_NAME, MEMORY_FEATURE_UNIT, power_draw, LABEL_NAME,
                         LABEL_UNIT, models_output_path, hostname)

    _plot_timeseries_data(timestamps, power_draw, LABEL_NAME, LABEL_UNIT, models_output_path, hostname)

    _plot_input_data_dependencies(dataset, models_output_path, hostname)


def _plot_timeseries_data(timestamps, feature, feature_name, feature_unit, models_output_path, hostname):
    # unify timestamps to count time from 00:00
    first_timestamp = timestamps[0]
    unified_timestamps = []
    for timestamp in timestamps:
        unified_timestamp = timestamp - first_timestamp
        unified_timestamps.append(unified_timestamp)

    # convert timestamps to datetime.datetime, then to matplotlib datenums to preserve seconds
    dates = [datetime.datetime.utcfromtimestamp(timestamp) for timestamp in unified_timestamps]
    date_nums = matplotlib.dates.date2num(dates)

    _create_plot(f'{feature_name} usage in time for host: {hostname}')

    plt.subplots_adjust(bottom=0.2)
    plt.xticks(rotation=90)

    # display more dates on x axis
    plt.locator_params(axis='x', tight=True, nbins=11)

    x_axis = plt.gca()
    date_formatter = matplotlib.dates.DateFormatter('%H:%M:%S')
    x_axis.xaxis.set_major_formatter(date_formatter)

    plt.plot(date_nums, feature)
    plt.xlabel(f'Elapsed test time [H:M:S]')
    plt.ylabel(f'{feature_name} [{feature_unit}]')
    plt.grid(True)

    figure_filename = f'{hostname}_timeseries_data_{feature_name}.png'
    _save_plot(models_output_path, hostname, figure_filename)


def _plot_input_data(input_feature, feature_name, feature_unit, label, label_name, label_unit, models_output_path,
                     hostname):
    _create_plot(f'Relationship between {feature_name} and {label_name} for host: {hostname}')

    plt.scatter(input_feature, label)
    plt.xlabel(f'{feature_name} [{feature_unit}]')
    plt.ylabel(f'{label_name} [{label_unit}]')
    plt.grid(True)

    figure_filename = f'{hostname}_input_data_{feature_name}.png'
    _save_plot(models_output_path, hostname, figure_filename)


def _plot_input_data_dependencies(dataset, models_output_path, hostname):
    columns = list(dataset.columns.values)
    plot = sns.pairplot(dataset[columns], diag_kind='kde')

    figure_directory_path = os.path.join(models_output_path, 'plots', hostname)
    create_directory(figure_directory_path)

    figure_filename = f'{hostname}_input_data_dependencies.png'
    figure_file_path = os.path.join(figure_directory_path, figure_filename)
    plot.figure.savefig(figure_file_path)


def plot_loss(history, models_output_path, hostname, model_name):
    _create_plot(f'Training loss for host: {hostname}')

    loss = history.history['loss']
    val_loss = history.history['val_loss']

    plt.plot(loss, label='loss')
    plt.plot(val_loss, label='val_loss')
    plt.xlabel('Epoch')
    plt.ylabel('Error (mean squared error)')
    plt.legend()
    plt.grid(True)

    figure_filename = f'{model_name}_loss.png'
    _save_plot(models_output_path, hostname, figure_filename)


def plot_predictions(labels, predicted_labels, models_output_path, hostname, model_name):
    _create_plot(f'True and predicted labels for host: {hostname} and model: {model_name}')
    # a = plt.axes(aspect='equal')

    plt.scatter(labels, predicted_labels)
    plt.xlabel(f'True values [{LABEL_NAME}]')
    plt.ylabel(f'Predictions [{LABEL_NAME}]')

    min_lim = min(min(labels), min(predicted_labels))
    max_lim = max(max(labels), max(predicted_labels))
    lim_margin = (max_lim - min_lim) * 0.1
    limits = [min_lim - lim_margin, max_lim + lim_margin]
    plt.xlim(limits)
    plt.ylim(limits)

    _ = plt.plot(labels, labels)

    figure_filename = f'{model_name}_predictions.png'
    _save_plot(models_output_path, hostname, figure_filename)


def plot_error_distribution(labels, predicted_labels, models_output_path, hostname, model_name):
    error = predicted_labels - labels

    _create_plot(f'Distribution of prediction error for host: {hostname} and model: {model_name}')

    plt.hist(error, bins=25)
    plt.xlabel(f'Prediction error of {LABEL_NAME}')
    _ = plt.ylabel('Count')

    figure_filename = f'{model_name}_error-distribution.png'
    _save_plot(models_output_path, hostname, figure_filename)


def _create_plot(figure_title):
    plt.figure()
    plt.title(figure_title)


def _save_plot(models_output_path, hostname, figure_filename):
    figure_directory_path = os.path.join(models_output_path, 'plots', hostname)
    create_directory(figure_directory_path)

    figure_file_path = os.path.join(figure_directory_path, figure_filename)

    plt.savefig(figure_file_path)
    plt.close()
