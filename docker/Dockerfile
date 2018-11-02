# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Use an official Go runtime as a parent image
FROM golang:1.11

# Install AMP Packager
RUN go get -v github.com/ampproject/amppackager/cmd/amppkg

# Seed the ocsp cache
WORKDIR /go/src/github.com/ampproject/amppackager/testdata/b1/
RUN mkdir -p /tmp/amppkg && touch ./index.txt && touch ./index.txt.attr && openssl ocsp -index ./index.txt -rsigner ca.ocsp.cert -rkey ca.privkey -CA ca.cert -ndays 7 -reqin ocspreq.der -respout /tmp/amppkg/ocsp

# Copy the binary and config file
WORKDIR /go/src/app
RUN cp /go/bin/amppkg .
COPY . .

# Make port 8080 available to the world outside this container. This
# port must match the AMP Packager port configured in the toml file.
EXPOSE 8080

# Start the AMP Packager
ENTRYPOINT ["amppkg"]

# Set default flags to run in development mode.
CMD ["-development", "-config=amppkg.example.toml"]

