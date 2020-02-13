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

  1. Visit the [Kubernetes Engine page](https://console.cloud.google.com/) in the Google Cloud Console.

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

  1. PROJECT_ID. This is your Google Cloud Project ID where you want your cluster to reside.

  2. COMPUTE_ENGINE_ZONE. Select a region where you want your app's computing resources located. See: [App Engine Locations](https://cloud.google.com/appengine/docs/locations)

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

## Optionally, run the teardown script (gcloud_down.sh)

  1. Go to the directory where you installed amppackager.
  2. cd deploy/gcloud
  3. ./gcloud_down.sh
  4. Wait for script to finish.

