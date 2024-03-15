from setuptools import setup, find_packages

setup(
    name='storage_service_grpc',
    version='0.0.4',
    packages=find_packages(),
    install_requires=[
        "grpcio",
        "grpcio-tools"
    ],
)