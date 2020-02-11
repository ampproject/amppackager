#!/bin/bash
gcloud container clusters delete amppackager-cluster --quiet --verbosity=none
gcloud compute disks delete amppackager-nfs-disk --quiet --verbosity=none
