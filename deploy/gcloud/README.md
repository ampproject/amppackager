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
  DigiCert has deprecated access to the old style ACME Directory URL from CertCentral UI in order to transition to using EAB/HMAC (which we will support in a future release). In the meantime, you can retrieve the old style ACME Directory URL using:

      curl --location --request POST 'https://www.digicert.com/services/v2/key/acme-url' --header 'X-DC-DEVKEY: $YOUR_DEV_KEY' --header 'Content-Type: application/json' --data-raw '{
      "name": "myTest",
      "product_name_id": "ssl_plus",
      "organization_id": "$YOUR_ORG_ID",
      "order_validity_days": "350",
      "validity_days": "90",
      "container_id": "$YOUR_CONTAINER_ID",
      "profile_option": "http_signed_exchange",
      "external_account_binding": false}'
      
   $YOUR_DEV_KEY can be found [here](https://www.digicert.com/secure/automation/api-keys/). Add with no restrictions.
   
   $YOUR_ORG_ID can be found [here](https://www.digicert.com/secure/organizations/).
   
   $YOUR_CONTAINER_ID can be generated using:
   
      curl -X GET https://www.digicert.com/services/v2/container \
         -H 'Content-Type: application/json' \
         -H 'X-DC-DEVKEY: $YOUR_DEV_KEY'

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

## Production Reverse Proxy Setup

For now, productionizing is a bit manual. The minimum steps are:

  1. Don't expose amppkg to the outside world. If possible, keep it on your internal network.
     In addition, use the loadBalancerSourceRanges above so that it only responds to the reverse
     proxy you set up. More information could be found here:
     1. [Kubernetes Firewall Configuration](https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/)
     2. [Subnet Calculator](https://mxtoolbox.com/subnetcalculator.aspx)
     3. [Understanding CIDR notation](https://www.digitalocean.com/community/tutorials/understanding-ip-addresses-subnets-and-cidr-notation-for-networking)
  2. Configure your TLS-serving frontend server to conditionally proxy to
     `amppkg`:
     1. If the URL starts with `/amppkg/`, forward the request unmodified.
     2. If the URL points to an AMP page and the `AMP-Cache-Transform` request
        header is present, rewrite the URL by prepending `/priv/doc` and forward
        the request.

        NOTE: If using nginx, prefer using `proxy_pass` with `$request_uri`,
        rather than using `rewrite`, as in [this PR](https://github.com/Warashi/try-amppackager/pull/3),
        to avoid percent-encoding issues.
     3. If at all possible, don't send URLs of non-AMP pages to `amppkg`; its
        [transforms](transformer/) may break non-AMP HTML.
     4. DO NOT forward `/priv/doc` requests; these URLs are meant to be
        generated by the frontend server only.
  3. For HTTP compliance, ensure the `Vary` header set to `AMP-Cache-Transform,
     Accept` for all URLs that point to an AMP page, irrespective of whether the
     response is HTML or SXG. (SXG responses that come from `amppkg` will have
     the appropriate `Vary` header set, so it may only be necessary to
     explicitly set the `Vary` header for HTML responses.)
  4. Get an SXG cert from your CA. It must use an EC key with the prime256v1
     algorithm, and it must have a [CanSignHttpExchanges
     extension](https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cross-origin-cert-req).
     One provider of SXG certs is [DigiCert](https://www.digicert.com/account/ietf/http-signed-exchange.php).
     You can fill in your ACME Directory URL using these [steps](https://docs.digicert.com/certificate-tools/Certificate-lifecycle-automation-index/acme-user-guide/acme-directory-urls-signed-http-exchange-certificates/).
  5. Every 90 days or sooner, renew your SXG cert (per
     [WICG/webpackage#383](https://github.com/WICG/webpackage/pull/383)) and
     restart amppkg (per
     [#93](https://github.com/ampproject/amppackager/issues/93)). This is only necessary if ACME
     cert auto-renewal is not enabled.
  6. Keep amppkg updated from either `releases` (the default branch, so `go get` works)
     about every ~2 months. The [wg-caching](https://github.com/ampproject/wg-caching)
     team will release a new version approximately this often. Soon after each
     release, Googlebot will increment the version it requests with
     `AMP-Cache-Transform`. Googlebot will only allow the latest 2-3 versions
     (details are still TBD), so an update is necessary but not immediately. If
     amppkg doesn't support the requested version range, it will fall back to
     serving unsigned AMP.

     The version of amppkg on the [Google Cloud Marketplace](https://cloud.google.com/marketplace/)
     will also be updated once a new release is published.

     To keep subscribed to releases, you can select "Releases only" from the
     "Watch" dropdown in GitHub, or use [various tools](https://stackoverflow.com/questions/9845655/how-do-i-get-notifications-for-commits-to-a-repository)
     to subscribe to the `releases` branch.

## Optionally, run the teardown script (gcloud_down.sh)

  1. Go to the directory where you installed amppackager.
  2. cd deploy/gcloud
  3. ./gcloud_down.sh
  4. Wait for script to finish.

