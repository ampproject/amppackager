#!/bin/bash
source $(dirname $0)/setup.sh

gcloud config set project $PROJECT_ID
gcloud config set compute/zone $COMPUTE_ENGINE_ZONE

gcloud container clusters delete amppackager-cluster --quiet --verbosity=none
gcloud compute disks delete amppackager-nfs-disk --quiet --verbosity=none
