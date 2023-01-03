import os


def create_directory(path):
    if not os.path.isdir(path):
        os.makedirs(path)
