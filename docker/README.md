## Overview

This is a base docker image that runs the AMP Packager in development
mode locally.

Stay tuned for more instructions on productionizing.

### Run locally in development mode

#### As-is

To run the AMP Packager with the fake test certificates that it ships
with (forwarding local port 8080 to the container's port 8080):

```sh
$ docker build -t amppackager .
$ docker run -p 8080:8080 amppackager
```

Refer to [these instructions](../README.md#test-your-config) on how to run an
end-to-end test.

#### Customizing

To use your own certificates, create your own Dockerfile using this
one as a base image. In your Dockerfile, you can copy in your custom
.toml file and certificates.

Write a Dockerfile something along the lines of:

```
FROM amppackager

WORKDIR /go/src/app

# Copy in your .toml and certs. Adjust this command as necessary for
# correct source and target destinations.
COPY . .

# Change the default flags to use your config
CMD [ "-development", "-config=<path_to_your_toml_on_docker>"]

```

Then build your Docker image and run it.
