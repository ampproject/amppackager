#!/bin/bash
# Build docker images for the AMP packager renewer, consumer and init.
# Renewer and consumer are the same binaries, passed different command line
# arguments. These images can be used for testing the AMP Packager deployer and
# may also be installed in the gcloud marketplace.
export PROJECT_ID="YOUR_GCLOUD_PROJECT_ID"
export AMP_PACKAGER_VERSION_TAG="0.0"

docker build -f Dockerfile.consumer -t gcr.io/${PROJECT_ID}/amppackager:${AMP_PACKAGER_VERSION_TAG} .
docker build -f Dockerfile.renewer -t gcr.io/${PROJECT_ID}/amppackager_renewer:${AMP_PACKAGER_VERSION_TAG} .
docker build -f Dockerfile.init -t gcr.io/${PROJECT_ID}/amppackager_init:${AMP_PACKAGER_VERSION_TAG} .

docker push gcr.io/${PROJECT_ID}/amppackager:${AMP_PACKAGER_VERSION_TAG}
docker push gcr.io/${PROJECT_ID}/amppackager_renewer:${AMP_PACKAGER_VERSION_TAG}
docker push gcr.io/${PROJECT_ID}/amppackager_init:${AMP_PACKAGER_VERSION_TAG}
