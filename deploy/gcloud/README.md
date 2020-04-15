# AMP Packager Google Cloud Deployment

This page shows you how to deploy a [Kubernetes](https://kubernetes.io/) cluster of
 [AMP Packager](https://github.com/ampproject/amppackager#amp-packager) in a Google Cloud environment:

  1. Setup your production or development environment in preparation for the cloud deployment.
  2. Modify the setup.sh script to enter all information relevant to your cloud
     deployment.
  3. Run the gcloud_up.sh to initiate the deployment.
  4. Run the gcloud_down.sh to tear down the deployment, if desired.

## Before you begin

These instructions are mostly lifted from the [Google Kubernetes Engine tutorial](https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app).  You may want to consult that page for further reference.  The following items are required before you can begin your deployment:

  1. Visit the [Kubernetes Engine page](https://console.cloud.google.com/kubernetes/list) in the Google Cloud Console.

  2. Create or select an existing [Google Cloud Project](https://cloud.google.com/appengine/docs/standard/nodejs/building-app/creating-project)

  3. Wait for the API and related services to be enabled. This can take several minutes.

  4. Make sure that billing is enabled for your Google Cloud project. [Learn how to confirm billing is enabled for your project](https://cloud.google.com/billing/docs/how-to/modify-project).

  5. Install the [Google Cloud SDK](https://cloud.google.com/sdk/docs/quickstarts), which includes the gcloud command-line tool.

  6. Using the gcloud command line tool, install the Kubernetes command-line tool. kubectl is used to communicate with Kubernetes, which is the cluster orchestration system of GKE clusters:

    gcloud components install kubectl

  7. Install [Docker Community Edition (CE)](https://docs.docker.com/install/) on your workstation. You will use this to build a container image for the application.

  8. Install the [Git source control](https://git-scm.com/downloads) tool to fetch the sample application from GitHub.

## Enter your project setup information into setup.sh

The following information is required to be entered into setup.sh:

  1. PROJECT_ID. This is your Google Cloud Project ID where you want your cluster to reside. Note that if your project ID is scoped by a domain, you need to replace the ':' with a '/' in the project id name. See: [Google Cloud domain scoped projects](https://cloud.google.com/container-registry/docs/overview#domain-scoped_projects).

  2. COMPUTE_ENGINE_ZONE. Select a region where you want your app's computing resources located. See: [Compute Zone Locations](https://console.cloud.google.com/compute/zones)

  3. AMP_PACKAGER_DOMAIN.  The domain you want to use for the [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  4. AMP_PACKAGER_COUNTRY. The country you want to use for the [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  5. AMP_PACKAGER_STATE. The state you want to use for the [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  6. AMP_PACKAGER_LOCALITY. The locality you want to use for the [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  7. AMP_PACKAGER_ORGANIZATION. The organization you want to use for the [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  8. ACME_EMAIL_ADDRESS. The email address you used for your [Digicert ACME Account](https://docs.digicert.com/manage-certificates/certificate-profile-options/get-your-signed-http-exchange-certificate/).

  9. ACME_DIRECTORY_URL.  The [ACME API Directory URL](https://docs.digicert.com/certificate-tools/acme-user-guide/acme-directory-urls-signed-http-exchange-certificates/).  Note that this URL is security sensitive so do not check in any files that contain this into Github.

The following information can be customized in setup.sh, but the default also works fine:

  1. AMP_PACKAGER_VERSION_TAG. The version tag attached to your docker image when it's built and uploaded into the [Container Registry](https://cloud.google.com/container-registry).

  2. AMP_PACKAGER_NUM_REPLICAS.  Number of AMP Packager replicas to run. Default is 2, you can scale up to 8 instances.

  3. AMP_PACKAGER_CERT_FILENAME. The name of the amppackager certificate.

  4. AMP_PACKAGER_CSR_FILENAME. The name of the amppackager [Certificate Signing Request](https://www.digicert.com/ecc-csr-creation-ssl-installation-apache.htm).

  5. AMP_PACKAGER_PRIV_KEY_FILENAME. The name of the amppackager private key (see this [article])(https://www.ericlight.com/using-ecdsa-certificates-with-lets-encrypt).

  Note that the cert, CSR and private key all go together. You may choose to
  manually create these and [request the certificate from Digicert by email](https://docs.digicert.com/manage-certificates/certificate-profile-options/get-your-signed-http-exchange-certificate/). If you chose to go this route, you may simply copy these files into your
  amppackager/deploy/gcloud/generated directory and they will not be auto-generated and
  fulfilled by amppackager using ACME. By default, these files will be
  automatically created given the information you supplied in setup.sh, and the
  certificate will be requested using ACME automatically.

## Run the deployment script (gcloud_up.sh)

  1. Go to the directory where you installed amppackager.
  2. cd deploy/gcloud
  3. ./gcloud_up.sh. The script may pause and prompt you for a response at certain points before it
     continues.
  4. Wait for script to finish.

## Make sure the deployment is up and ready

  1. Check that all the components of the deployment are up in the Kubernetes
     page. If everything is alright, everything should have a green check mark
     in the console. Check everything under Cluster, Workloads and "Service &
     Ingress".

     Clusters will have "amppackager-cluster".
     Workloads will have "amppackager-cert-renewer-pd",
      "amppackager-nfs-server", and user-specified number of replicas of
      "amppackager-cert-consumer-deployment".
     Service & Ingress will have "amppackager-nfs-service and
      amppackager-service".

  2. Issue curl command to check on the health of the amppackager. It should
     return "ok". The amppackager service IP address could be using "kubectl get
     service, it will be under EXTERNAL-IP column.

    curl http://AMP_PACKAGER_SERVICE_IP_ADDRESS:AMP_PACKAGER_SERVICE_PORT/healthz

  3. Issue curl command to test if you can download the sample signed exchange
     in your domain ($AMP_PACKAGER_DOMAIN in setup.sh). 

    curl -H 'Accept: application/signed-exchange;v=b3' -H 'AMP-Cache-Transform: google;v="1..2"' -i http://AMP_PACKAGER_SERVICE_IP_ADDRESS:AMP_PACKAGER_SERVICE_PORT/priv/doc/https://$AMP_PACKAGER_DOMAIN/

  4. After you finish testing, lock down the amppackager service so that it's
     only visible to your frontend server.  You do this by modifying the
     following section of amppackage_service.yaml:

      loadBalancerSourceRanges:
      - YOUR_FRONTEND_SERVER_IP_ADDRESS_HERE in CIDR format. CIDR is explained
        in the comments before this section in the yaml file.

  5. After you make the modifications in step 4, issue the following command to
     apply changes to you cluster:

  kubectl apply -f amppackager_service.yaml

## Optionally, run the teardown script (gcloud_down.sh)

  1. Go to the directory where you installed amppackager.
  2. cd deploy/gcloud
  3. ./gcloud_down.sh
  4. Wait for script to finish.

