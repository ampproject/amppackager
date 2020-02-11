# git clone https://github.com/ampproject/amppackager.git
# cd amppackager/deploy/gcloud
# To start: ./gcloud-up.sh
# To shutdown: ./gcloud-down.sh

# PROJECT_ID = gcloud project id
export PROJECT_ID="azei-package-test"
export COMPUTE_ENGINE_ZONE="us-west1-a"

# AMP_PACKAGER_VERSION_TAG
export AMP_PACKAGER_VERSION_TAG="latest"

# Build docker images for the Amppackager renewer and consumer, if necessary.
# Renewer and consumer are the same binaries, passed difference command line
# arguments.
docker build -f Dockerfile.consumer -t gcr.io/${PROJECT_ID}/amppackager:${AMP_PACKAGER_VERSION_TAG} .
docker build -f Dockerfile.renewer -t gcr.io/${PROJECT_ID}/amppackager_renewer:${AMP_PACKAGER_VERSION_TAG} .

# To check that it succeeded, list the images that got built.
docker images

# https://cloud.google.com/container-registry/docs/advanced-authentication#gcloud-helper
gcloud auth configure-docker

# Push the docker image into the cloud container registry.
# See: https://cloud.google.com/container-registry/docs/overview
# See: https://cloud.google.com/container-registry/docs/pushing-and-pulling
docker push gcr.io/${PROJECT_ID}/amppackager:${AMP_PACKAGER_VERSION_TAG}
docker push gcr.io/${PROJECT_ID}/amppackager_renewer:${AMP_PACKAGER_VERSION_TAG}

gcloud config set project $PROJECT_ID
gcloud config set compute/zone $COMPUTE_ENGINE_ZONE

# Allow 10 nodes maximum for this cluster.
#gcloud container clusters create amppackager-cluster --tags allow-8080 --num-nodes=10
gcloud container clusters create amppackager-cluster --num-nodes=10

# Setup your credentials
# https://cloud.google.com/sdk/gcloud/reference/container/clusters/get-credentials
gcloud container clusters get-credentials amppackager-cluster

# Create the NFS disk for RW sharing amongst the kubernetes deplouments
# cert-renewer and cert-consumer.
# https://medium.com/@Sushil_Kumar/readwritemany-persistent-volumes-in-google-kubernetes-engine-a0b93e203180
gcloud compute disks create --size=10GB --zone=${COMPUTE_ENGINE_ZONE} amppackager-nfs-disk
kubectl apply -f nfs-renewer-pvc.yaml
kubectl apply -f nfs-consumer-pvc.yaml
kubectl apply -f nfs-clusterip-service.yaml
kubectl apply -f nfs-server-deployment.yaml
export AMPPACKAGER_NFS_SERVER=$(kubectl get pods | grep amppackager-nfs | awk '{print $1}')

# Sleep for a few minutes, waiting for NFS disk to be deployed.
sleep 4m

# This assumes current working directory is amppackager/docker/gcloud
# default is the default namespace for the gcloud project
kubectl cp www default/"$AMPPACKAGER_NFS_SERVER":/exports/
kubectl cp amppkg_consumer.toml default/"$AMPPACKAGER_NFS_SERVER":/exports
kubectl cp amppkg_renewer.toml default/"$AMPPACKAGER_NFS_SERVER":/exports
kubectl cp amppkg.cert default/"$AMPPACKAGER_NFS_SERVER":/exports
kubectl cp amppkg.privkey default/"$AMPPACKAGER_NFS_SERVER":/exports
kubectl cp amppkg.csr  default/"$AMPPACKAGER_NFS_SERVER":/exports

kubectl apply -f amppackager-cert-renewer.yaml
kubectl apply -f amppackager-cert-consumer.yaml
kubectl apply -f amppackager-service.yaml

# List the service that got started.
kubectl get service
