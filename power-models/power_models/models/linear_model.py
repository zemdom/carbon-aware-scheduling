import os.path

import numpy as np
import tensorflow as tf
from tensorflow import keras

from constants import PIECEWISE_LINEAR_CPU_MODELS_COUNT, VALIDATION_SPLIT

from utils.plots import plot_loss, plot_predictions, plot_error_distribution


# reference: https://www.tensorflow.org/tutorials/keras/regression
def create_model(train_features, use_memory_feature):
    # add normalization layers, as CPU and memory data have different range
    normalizer = keras.layers.Normalization(input_shape=[2 if use_memory_feature else 1], axis=-1)
    normalizer.adapt(np.array(train_features))

    model = keras.Sequential([
        normalizer,
        keras.layers.Dense(units=1)
    ])

    model.compile(
        optimizer=keras.optimizers.SGD(),
        loss=keras.losses.MeanSquaredError()
    )

    model.summary()

    return model


def train_model(model: keras.Sequential, train_features, train_labels, models_output_path, hostname, model_index,
                epochs, create_piecewise_linear_cpu_model):
    history = model.fit(
        train_features,
        train_labels,
        epochs=epochs,
        validation_split=VALIDATION_SPLIT
    )

    model_name = _generate_model_name(model_index, hostname, create_piecewise_linear_cpu_model)
    plot_loss(history, models_output_path, hostname, model_name)

    return model


def use_model(model: keras.Sequential, test_features, test_labels, models_output_path, hostname, model_index,
              create_piecewise_linear_cpu_model):
    predicted_labels = model.predict(test_features).flatten()

    model_name = _generate_model_name(model_index, hostname, create_piecewise_linear_cpu_model)
    plot_predictions(test_labels, predicted_labels, models_output_path, hostname, model_name)
    plot_error_distribution(test_labels, predicted_labels, models_output_path, hostname, model_name)

    return predicted_labels[0]


def save_model(model: keras.Sequential, model_output_path, hostname, model_index, create_piecewise_linear_cpu_model):
    model_name = _generate_model_name(model_index, hostname, create_piecewise_linear_cpu_model)

    model_full_output_path = os.path.join(model_output_path, hostname, model_name)
    tf.saved_model.save(model, model_full_output_path)

    loaded_model = tf.saved_model.load(model_full_output_path)
    input_tensor = loaded_model.signatures['serving_default'].inputs[0]

    with open(os.path.join(model_output_path, hostname, f'{model_name}_input-tensor-info.log'), 'w') as output_file:
        input_tensor_shape = str(input_tensor.shape)
        input_tensor_shape = input_tensor_shape.replace('None', '1')
        input_tensor_shape = input_tensor_shape.replace(' ', '')
        output_file.write(input_tensor_shape)
        output_file.write('\n')
        output_file.write(input_tensor.dtype.name)

    # delete tensorflow graph to allow next models to use the same SignatureDefs (names of layers)
    keras.backend.clear_session()


def _generate_model_name(model_index, hostname, create_piecewise_linear_cpu_model):
    piecewise_linear_cpu_model = f'{model_index + 1}-of-{PIECEWISE_LINEAR_CPU_MODELS_COUNT}' \
        if create_piecewise_linear_cpu_model else ''
    model_name = f'{hostname}_model{piecewise_linear_cpu_model}'

    return model_name
